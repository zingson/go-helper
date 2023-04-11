package ccblife_auth

import (
	"encoding/json"
	"net/url"
)

// 授权用户参数
// BGCOLOR=&userid=YSM202112210902914&mobile=13611703040&cityid=330200&userCityId=310000&orderid=&PLATFLOWNO=0000A2UNK1640064349000416&openid=&lgt=121.45578&ltt=31.275461&Usr_Name=&USERID=YSM202112210902914&MOBILE=13611703040&CITYID=330200&USERCITYID=310000&ORDERID=&OPENID=&LGT=121.45578&LTT=31.275461&timestamp=20220127114219&TIMESTAMP=1643254939767

type CcbParamSJ struct {
	Userid     string `json:"userid"`     // 建行生活用户id
	Mobile     string `json:"mobile"`     // 手机号
	Cityid     string `json:"cityid"`     // 选择的城市编号
	UserCityId string `json:"userCityId"` // 用户定位的城市编号
	Timestamp  string `json:"timestamp"`  // 跳转场景方时间戳
	Platflowno string `json:"platflowno"` //PLATFLOWNO为登录校验流水号

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
		Cityid:     values.Get("cityid"),
		UserCityId: values.Get("userCityId"),
		Platflowno: values.Get("PLATFLOWNO"),
		Timestamp:  values.Get("timestamp"),
	}
	return
}
