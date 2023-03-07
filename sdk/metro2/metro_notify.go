package metro2

import "net/http"

//ResNotify 进出站通知
func ResNotify(config *Config, rid string, request *http.Request, response http.ResponseWriter, f func(v *NotifyData) (err error)) {
	httpCallback(config, rid, request, response, func(reqBody string) (resData any, err error) {
		v, err := jsonParse[*NotifyData](reqBody)
		if err != nil {
			return
		}
		err = f(v)
		if err != nil {
			return
		}
		resData = make(map[string]string)
		return
	})
}

type NotifyData struct {
	GuidVoucher     string `json:"GuidVoucher"`     // 票卡编号
	Times           int64  `json:"Times"`           // 次数
	RemainTimes     int64  `json:"RemainTimes"`     // 剩余次数
	BeginStatin     string `json:"BeginStatin"`     // 进站
	BeginStatinTime string `json:"BeginStatinTime"` // 进站时间
	EndStatin       string `json:"EndStatin"`       // 出站
	EndStatinTime   string `json:"EndStatinTime"`   // 出站时间
}
