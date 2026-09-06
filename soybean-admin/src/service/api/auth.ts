import { request } from '../request';

/**
 * Login step 1: verify username + password.
 *
 * Never returns a session token directly - the backend always responds with
 * a short-lived TOTP challenge (either "enroll" for first-time setup or
 * "verify" for an already-enrolled admin). Call {@link fetchLoginVerify}
 * with the returned challenge and an authenticator code to complete login.
 */
export function fetchLogin(model: Api.Form.LoginForm) {
  return request<Api.Auth.LoginChallenge>({
    url: '/auth/login',
    method: 'post',
    data: model
  });
}

/** Login step 2: complete a TOTP challenge and receive a session token. */
export function fetchLoginVerify(params: { challenge: string; code: string }) {
  return request<Api.Auth.LoginToken>({
    url: '/auth/login/verify',
    method: 'post',
    data: params
  });
}

/** Get user info */
export function fetchGetUserInfo() {
  return request<Api.Auth.UserInfo>({ url: '/userinfo' });
}
