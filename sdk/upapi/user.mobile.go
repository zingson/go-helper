package upapi

// UserMobile 取用户手机号码
func UserMobile(conf *Config, accessToken, openId string, backendToken func(config *Config) string) (r *UserMobileResult, err error) {
	bm := NewBodyMap()
	bm.Set("appId", conf.Appid)
	bm.Set("backendToken", backendToken(conf))
	bm.Set("accessToken", accessToken)
	bm.Set("openId", openId)

	err = Call(conf, "/user.mobile", bm, &r)
	// 解密密文手机号
	r.Mobile, err = Decode3DES(conf.SymmetricKey, r.Mobile)
	if err != nil {
		return
	}
	return
}

type UserMobileResult struct {
	Mobile string `json:"mobile"`
}
