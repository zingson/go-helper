package metro

import "errors"

// TicketOpen 地铁票开通
func TicketOpen(cfg *Config, mobile, productCode, outOrderNo string) (ticketCode *TicketCode, err error) {
	userId, err := AuthByMobile(cfg, mobile)
	if err != nil {
		return
	}
	data, err := MonthlyTicketOpen(cfg, userId, productCode, outOrderNo)
	if err != nil {
		return
	}
	if data.Tickets == nil || len(data.Tickets) == 0 {
		err = errors.New("METRO:票开通未返回票码")
		return
	}
	tk := data.Tickets[0]

	ticketCode = &TicketCode{
		OutOrderNo:      outOrderNo,
		OrderNo:         data.OrderNo,
		TicketCode:      tk.TicketCode,
		TicketTimes:     tk.TicketTimes,
		TicketStartTime: tk.TicketStartTime,
		TicketEndTime:   tk.TicketEndTime,
	}
	return
}

type TicketCode struct {
	OutOrderNo      string `json:"outOrderNo"`      // 接入方系统订单号
	OrderNo         string `json:"orderNo"`         // 票务平台订单号
	TicketCode      string `json:"ticketCode"`      //月票编号
	TicketTimes     int    `json:"ticketTimes"`     //月票次数
	TicketStartTime string `json:"ticketStartTime"` //月票有效起始时间，yyyy-MM-ddHH:mm:ss
	TicketEndTime   string `json:"ticketEndTime"`   //月票有效截止时间，yyyy-MM-ddHH:mm:ss
}
