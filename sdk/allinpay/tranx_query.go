package allinpay

import (
	"errors"
)

//Query 交易查询 ,reqsn = 商户订单号 trxid=平台交易流水  reqsn和trxid必填其一 建议:商户如果同时拥有trxid和reqsn,优先使用trxid
//文档： https://aipboss.allinpay.com/know/devhelp/main.php?pid=15#mid=93
func Query(conf *Config, reqsn, trxid string) (result *QueryResult, err error) {
	var bm = make(map[string]string)
	bm["trxid"] = trxid
	bm["reqsn"] = reqsn

	err = PostForm(conf, "/apiweb/unitorder/query", bm, &result)
	if err != nil {
		return
	}
	if result.Trxstatus != TRX_0000 {
		err = errors.New(string(result.Trxstatus) + ":" + result.Errmsg)
		return
	}
	return
}

type QueryResult struct {
	Trxid     string    `json:"trxid"`
	Chnltrxid string    `json:"chnltrxid"`
	Reqsn     string    `json:"reqsn"`
	Trxcode   string    `json:"trxcode"`
	Trxamt    string    `json:"trxamt"`
	Trxstatus TrxStatus `json:"trxstatus"`
	Errmsg    string    `json:"errmsg"` //错误原因
	Acct      string    `json:"acct"`
	Fintime   string    `json:"fintime"`
	Cmid      string    `json:"cmid"`
	Chnlid    string    `json:"chnlid"`
	Initamt   string    `json:"initamt"`
	Fee       string    `json:"fee"`
	Chnldata  string    `json:"chnldata"`
}

//GetTrxamt 交易金额
func (m *QueryResult) GetTrxamt() int64 {
	return ParseAmt(m.Trxamt)
}

//GetFee 手续费
func (m *QueryResult) GetFee() int64 {
	return ParseAmt(m.Fee)
}
