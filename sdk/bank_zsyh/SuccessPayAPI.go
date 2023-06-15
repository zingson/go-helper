package bank_zsyh

// SuccessPayApi 成功支付结果通知API
// 一网通支付API中的payNoticeUrl参数
// 文档：http://openhome.cmbchina.com/PayNew/pay/doc/cell/H5/SuccessPayAPI
func SuccessPayApi(conf *Config, body []byte) (noticeData SuccessPayData, err error) {
	return ParseNotice[SuccessPayData](conf, body)
}

type SuccessPayData struct {
	DateTime       string `json:"dateTime"`       //
	NoticeUrl      string `json:"noticeUrl"`      //
	HttpMethod     string `json:"httpMethod"`     //
	BranchNo       string `json:"branchNo"`       //
	MerchantNo     string `json:"merchantNo"`     //
	NoticeType     string `json:"noticeType"`     //本接口固定为：“BKPAYRTN”
	NoticeSerialNo string `json:"noticeSerialNo"` //银行通知序号,订单日期+订单号
	Date           string `json:"date"`           //商户订单日期 格式：yyyyMMdd
	OrderNo        string `json:"orderNo"`        //商户订单号
	Amount         string `json:"amount"`         //订单金额,单位元，格式：XXXX.XX
	BankDate       string `json:"bankDate"`       //
	BankSerialNo   string `json:"bankSerialNo"`   //
	DiscountFlag   string `json:"discountFlag"`   //
	DiscountAmount string `json:"discountAmount"` //
	MerchantPara   string `json:"merchantPara"`   //
	CardType       string `json:"cardType"`       //
	UniqueUserID   string `json:"uniqueUserID"`   //联系招行配置控制参数后返回（默认不返回） 支付用户ID，由身份证+姓名生成
	ExpandUserID   string `json:"expandUserID"`   //联系招行配置控制参数后返回（默认不返回） 一网通用户ID，由一网通ID生成（招行一网通用户的身份标识）
}
