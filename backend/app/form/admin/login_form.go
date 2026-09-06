package admin

type LoginForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginVerifyForm struct {
	Challenge string `json:"challenge"`
	Code      string `json:"code"`
}
