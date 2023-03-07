package metro2

import (
	"errors"
	"fmt"
)

//Register 6.1.2.1第三方用户注册
func Register(config *Config, rid string, GuidUser, MobilePhone string) (err error) {
	bm := make(map[string]string)
	bm["GuidUser"] = GuidUser
	bm["MobilePhone"], err = EncodeCBC(MobilePhone, config.AesSecret, config.AesIv)
	if err != nil {
		return
	}

	resData, err := HttpPost(config, rid, "/extp/tp/user/register", jsonStringify(bm))
	if err != nil {
		return
	}

	regRes, err := jsonParse[*RegisterResponse](resData)
	if err != nil {
		return
	}
	if regRes.Code != 200 {
		err = errors.New(fmt.Sprintf("METRO2_%d:%s", regRes.Code, regRes.Message))
		return
	}
	return
}

type RegisterResponse struct {
	GuidRequest string `json:"GuidRequest"`
	Code        int64  `json:"Code"`
	Message     string `json:"Message"`
}
