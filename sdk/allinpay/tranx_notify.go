package allinpay

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

// Notify 交易结果通知  文档：https://aipboss.allinpay.com/know/devhelp/main.php?pid=15#mid=94
func Notify(req *http.Request, config func(cusid string) *Config, resolve func(tranx *TranxNotify) error) (resBody string) {
	var (
		rid     = Rand32()
		bmilli  = time.Now().UnixMilli()
		err     error
		reqBody string
	)
	defer func() {
		if err != nil {
			resBody = err.Error() + " ,rid=" + rid
		}
		millisecond := fmt.Sprintf("%d", time.Now().UnixMilli()-bmilli)
		logrus.WithField("rid", rid).WithField("millisecond", millisecond).Infof("sdk allinpay 收银宝支付通知 请求URI：%s  请求报文：%s  响应报文：%s  | %sms", req.RequestURI, reqBody, resBody, millisecond)
	}()

	err = req.ParseForm()
	if err != nil {
		return
	}
	var bm = make(map[string]string)
	for k, v := range req.PostForm {
		bm[k] = v[0]
	}
	bmBytes, err := json.Marshal(bm)
	if err != nil {
		return
	}
	reqBody = string(bmBytes) // 请求参数转为JSON格式

	var tranx *TranxNotify
	err = json.Unmarshal(bmBytes, &tranx)
	if err != nil {
		err = errors.New("通知报文解析失败 " + err.Error())
		return
	}
	if tranx == nil {
		err = errors.New("通知报文解析失败")
		return
	}

	// 交易成功状态判断
	if tranx.Trxstatus != TRX_0000 {
		err = errors.New(string(tranx.Trxstatus) + ":未完成")
		return
	}

	// 读配置
	cfg := config(tranx.Cusid)
	if cfg == nil || cfg.PubKey == "" {
		err = errors.New("未读取到接口配置 cusid:" + tranx.Cusid)
		return
	}
	// 签名验证
	err = RsaVerify(cfg, bm)
	if err != nil {
		return
	}

	// 业务处理
	err = resolve(tranx)
	if err != nil {
		return
	}

	resBody = "success"
	return
}

type TranxNotify struct {
	Appid       string    `json:"appid"`       //收银宝APPID
	Trxcode     string    `json:"trxcode"`     //交易类型
	Trxid       string    `json:"trxid"`       //**通联收银宝交易流水号
	Trxamt      string    `json:"trxamt"`      //**交易金额 单位分
	Trxdate     string    `json:"trxdate"`     //yyyymmdd
	Paytime     string    `json:"paytime"`     //yyyymmddhhmmss
	Chnltrxid   string    `json:"chnltrxid"`   //渠道流水号  如支付宝，微信平台订单号
	Trxstatus   TrxStatus `json:"trxstatus"`   //**交易结果码
	Cusid       string    `json:"cusid"`       //**商户编号
	Cusorderid  string    `json:"cusorderid"`  //**统一下单对应的reqsn订单号
	Acct        string    `json:"acct"`        //卡号，可为空
	Fee         string    `json:"fee"`         //手续费，单位分
	Signtype    string    `json:"signtype"`    //前面类型 MD5或RSA。为空默认MD5
	Accttype    string    `json:"accttype"`    //借贷标识 00-借记卡  02-信用卡  99-其他（花呗/余额等）
	Chnldata    string    `json:"chnldata"`    //渠道信息   仅返回云闪付/微信/支付宝的渠道信息
	Srctrxid    string    `json:"srctrxid"`    //原交易流水 通联原交易流水，冲正撤销交易本字段不为空
	Trxreserved string    `json:"trxreserved"` //交易备注
}

//GetTrxamt 交易金额
func (m *TranxNotify) GetTrxamt() int64 {
	return ParseAmt(m.Trxamt)
}

//GetFee 手续费
func (m *TranxNotify) GetFee() int64 {
	return ParseAmt(m.Fee)
}

func (m *TranxNotify) Json() string {
	b, _ := json.Marshal(m)
	return string(b)
}
