package admin

import (
	"rustdesk-api-server-pro/app/form/admin"
	"rustdesk-api-server-pro/app/model"
	"rustdesk-api-server-pro/config"
	"rustdesk-api-server-pro/util"
	"strconv"
	"time"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/pquerna/otp/totp"
)

type AuthController struct {
	basicController
	Cfg *config.ServerConfig
}

// challengeTTL bounds how long an admin has to complete the TOTP step
// after a successful password check before the challenge expires.
const challengeTTL = 5 * time.Minute

// PostAuthLogin verifies username + password and issues a short-lived TOTP
// challenge. It never returns a session token directly: every admin login
// must complete PostAuthLoginVerify with a valid authenticator code, whether
// enrolling for the first time or confirming an already-enrolled secret.
func (c *AuthController) PostAuthLogin() mvc.Result {
	var loginForm admin.LoginForm
	if err := c.Ctx.ReadJSON(&loginForm); err != nil {
		return c.Error(nil, err.Error())
	}

	var user model.User
	get, err := c.Db.Where("username = ? and is_admin = 1", loginForm.Username).Get(&user)
	if err != nil {
		return c.Error(nil, err.Error())
	}
	if !get {
		return c.Error(nil, "UserNotExists")
	}
	if !util.PasswordVerify(loginForm.Password, user.Password) {
		return c.Error(nil, "UsernameOrPasswordError")
	}

	verifyCode := &model.VerifyCode{
		UserId:  user.Id,
		Type:    model.VC_TYPE_ADMIN_TFA,
		Uuid:    util.GetUUID(),
		Status:  model.VC_STATUS_UNUSED,
		Expired: time.Now().Add(challengeTTL),
	}

	if user.TwoFactorAuthSecret == "" {
		key, err := totp.Generate(totp.GenerateOpts{
			Issuer:      config.OPTIssuer,
			AccountName: user.Username,
		})
		if err != nil {
			return c.Error(nil, err.Error())
		}
		// Stash the pending secret on the challenge row; it is only
		// persisted onto the user once they prove possession of it below.
		verifyCode.Code = key.Secret()
		if _, err := c.Db.Insert(verifyCode); err != nil {
			return c.Error(nil, err.Error())
		}
		return c.Success(iris.Map{
			"stage":       "enroll",
			"challenge":   verifyCode.Uuid,
			"secret":      key.Secret(),
			"otpauth_url": key.URL(),
		}, "ok")
	}

	if _, err := c.Db.Insert(verifyCode); err != nil {
		return c.Error(nil, err.Error())
	}
	return c.Success(iris.Map{
		"stage":     "verify",
		"challenge": verifyCode.Uuid,
	}, "ok")
}

func (c *AuthController) PostAuthLoginVerify() mvc.Result {
	var form admin.LoginVerifyForm
	if err := c.Ctx.ReadJSON(&form); err != nil {
		return c.Error(nil, err.Error())
	}

	var verifyCode model.VerifyCode
	get, err := c.Db.Where("type = ? and uuid = ? and status = ?", model.VC_TYPE_ADMIN_TFA, form.Challenge, model.VC_STATUS_UNUSED).Get(&verifyCode)
	if err != nil {
		return c.Error(nil, err.Error())
	}
	if !get || verifyCode.Expired.Before(time.Now()) {
		return c.Error(nil, "TFA_Validate_Err")
	}

	var user model.User
	get, err = c.Db.Where("id = ? and is_admin = 1", verifyCode.UserId).Get(&user)
	if err != nil {
		return c.Error(nil, err.Error())
	}
	if !get {
		return c.Error(nil, "UserNotExists")
	}

	secret, enrolling := resolveTfaSecret(user.TwoFactorAuthSecret, verifyCode.Code)
	if secret == "" || !totp.Validate(form.Code, secret) {
		return c.Error(nil, "TFA_Validate_Err")
	}

	if _, err := c.Db.ID(verifyCode.Id).Cols("status").Update(&model.VerifyCode{Status: model.VC_STATUS_USED}); err != nil {
		return c.Error(nil, err.Error())
	}

	if enrolling {
		if _, err := c.Db.ID(user.Id).Cols("tfa_secret").Update(&model.User{TwoFactorAuthSecret: secret}); err != nil {
			return c.Error(nil, err.Error())
		}
	}

	// make other tokens expired
	_, _ = c.Db.Where("user_id = ? and status = 1 and is_admin = 1", user.Id).Cols("status").Update(&model.AuthToken{
		Status: 0,
	})

	signStr := strconv.Itoa(user.Id) + user.Username + time.Now().String()
	token := util.HmacSha256(signStr, c.Cfg.SignKey)
	expired := 2 * time.Hour // 2 hours

	authToken := &model.AuthToken{
		UserId:  user.Id,
		Token:   token,
		Expired: time.Now().Add(expired),
		IsAdmin: true,
		Status:  1,
	}

	if _, err := c.Db.Insert(authToken); err != nil {
		return c.Error(nil, err.Error())
	}

	return c.Success(iris.Map{
		"token": token,
	}, "ok")
}

// resolveTfaSecret decides which TOTP secret a submitted code must be
// validated against: the user's already-enrolled secret, or the pending
// secret stashed on the challenge row while enrollment is unconfirmed.
// It never mutates state - callers only persist pendingSecret once the
// submitted code has been validated against it.
func resolveTfaSecret(userSecret, pendingSecret string) (secret string, enrolling bool) {
	if userSecret != "" {
		return userSecret, false
	}
	return pendingSecret, true
}
