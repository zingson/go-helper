package upapi

//AccessToken 调用云闪付开放平台接口，通过CODE获取accessToken和openId
func AccessToken(conf *Config, code string, backendToken func(config *Config) string) (r *TokenResult, err error) {
	bm := NewBodyMap()
	bm.Set("appId", conf.Appid)
	bm.Set("backendToken", backendToken(conf))
	bm.Set("code", code)
	bm.Set("grantType", "authorization_code")

	err = Call(conf, "/token", bm, &r)
	return
}

type TokenResult struct {
	AccessToken string `json:"accessToken"`
	ExpiresIn   int64  `json:"expiresIn"`
	OpenId      string `json:"openId"`
	Scope       string `json:"scope"`
}
