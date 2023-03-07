package ccblife_auth

import (
	"encoding/json"
	"net/url"
)

// 授权用户参数
// BGCOLOR=&userid=YSM202112210902914&mobile=13611703040&cityid=330200&userCityId=310000&orderid=&PLATFLOWNO=0000A2UNK1640064349000416&openid=&lgt=121.45578&ltt=31.275461&Usr_Name=&USERID=YSM202112210902914&MOBILE=13611703040&CITYID=330200&USERCITYID=310000&ORDERID=&OPENID=&LGT=121.45578&LTT=31.275461&timestamp=20220127114219&TIMESTAMP=1643254939767

type CcbParamSJ struct {
	Userid     string `json:"userid"`     // 用户编号
	Mobile     string `json:"mobile"`     // 手机号
	Lgt        string `json:"lgt"`        //
	Ltt        string `json:"ltt"`        //
	Cityid     string `json:"cityid"`     //
	UserCityId string `json:"userCityId"` //
}

func (c *CcbParamSJ) Json() string {
	b, _ := json.Marshal(c)
	return string(b)
}

// ParseCcbParamSJ 解析建行生活授权用户参数
func ParseCcbParamSJ(conf *Config, v string) (param *CcbParamSJ, err error) {
	v, err = RsaDecode(v, conf.PriKey)
	if err != nil {
		return
	}

	values, err := url.ParseQuery(v)
	if err != nil {
		return
	}
	param = &CcbParamSJ{
		Userid:     values.Get("userid"),
		Mobile:     values.Get("mobile"),
		Lgt:        values.Get("lgt"),
		Ltt:        values.Get("ltt"),
		Cityid:     values.Get("cityid"),
		UserCityId: values.Get("userCityId"),
	}
	return
}
