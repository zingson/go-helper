package metro2

import (
	"errors"
	"fmt"
)

//Assign 6.1.2.2票卡领取
func Assign(config *Config, rid string, TradeNo, GuidUser, MobilePhone, GuidActivity string) (voucher *AssignResponseDataVoucher, err error) {
	// 每次发卡前做注册
	err = Register(config, rid, GuidUser, MobilePhone)
	if err != nil {
		return
	}

	bm := make(map[string]string)
	bm["TradeNo"] = TradeNo
	bm["GuidActivity"] = GuidActivity
	bm["GuidUser"] = GuidUser
	bm["MobilePhone"], err = EncodeCBC(MobilePhone, config.AesSecret, config.AesIv)
	if err != nil {
		return
	}

	resData, err := HttpPost(config, rid, "/extp/application/voucher/assign", jsonStringify(bm))
	if err != nil {
		return
	}

	assignRes, err := jsonParse[*AssignResponse](resData)
	if err != nil {
		return
	}
	if assignRes.Code != 200 {
		err = errors.New(fmt.Sprintf("METRO2_%d:%s", assignRes.Code, assignRes.Message))
		return
	}
	if assignRes.Data.Vouchers == nil || len(assignRes.Data.Vouchers) == 0 {
		err = errors.New("METR02_ERR:票卡领取响应报文数组读取失败")
		return
	}
	voucher = assignRes.Data.Vouchers[0]
	return
}

// {"GuidRequest":"","Code":200,"Message":"","Data":{"Vouchers":[{"GuidVoucher":"d35b70830a154b85abb1e65f5153bb82","Times":1,"TimeBegin":"2022-09-21 16:58:55","TimeEnd":"2022-09-21 23:59:59"}]}}

type AssignResponse struct {
	GuidRequest string              `json:"GuidRequest"`
	Code        int64               `json:"Code"`
	Message     string              `json:"Message"`
	Data        *AssignResponseData `json:"Data"`
}

type AssignResponseData struct {
	Vouchers []*AssignResponseDataVoucher `json:"Vouchers"`
}

type AssignResponseDataVoucher struct {
	GuidVoucher string `json:"GuidVoucher"` // 票卡编号
	Times       int64  `json:"Times"`       // 次数
	TimeBegin   string `json:"TimeBegin"`   // 有效期开始时间
	TimeEnd     string `json:"TimeEnd"`     // 有效期截止时间
}
