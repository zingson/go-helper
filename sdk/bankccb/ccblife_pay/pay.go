package ccblife_pay

import (
	"crypto/md5"
	"fmt"
	"time"
)

// 建行生活APP支付

/*
字段名 中文名 类型 是否必输 备注
MERCHANTID 商户代码 CHAR(15) Y 由建行统一分配
POSID 柜台代码 CHAR(9) Y 由建行统一分配
BRANCHID 分行代码 CHAR(9) Y 由建行统一分配
POSID19 商户19位终端号 CHAR(19) N 由建行统一分配，使用微信支付则必传，仅作为参数传递，不参与MAC校验
ORDERID 订单号 CHAR(30) Y 由商户提供，最长30位 PAYMENT 付款金额 NUMBER(16,2) Y 由商户提供，最长30位
PAYMENT 付款金额 NUMBER(16,2) Y 由商户提供，最长30位
CURCODE 币种 CHAR(2) Y 缺省为01－人民币（只支持人民币支付）
TXCODE 交易码 CHAR(6) Y 由建行统一分配为520100
REMARK1 备注1 CHAR(30) Y 网银不处理，直接传到城综网，该字段只支持送数字和英文
REMARK2 备注2 CHAR(30) Y 上送YS开头的服务方编号
TYPE 接口类型 CHAR(1) Y 1- 防钓鱼接口
GATEWAY 网关类型 CHAR(100) Y 默认送0
CLIENTIP 客户端IP CHAR(40) N 客户在商户系统中的IP
REGINFO 客户注册信息 CHAR(256) N 客户在商户系统中注册的信息，中文需使用escape编码
PROINFO 商品信息 CHAR(256) N 客户购买的商品，中文需使用escape编码
REFERER 商户URL CHAR(100) N 商户送空值即可
THIRDAPPINFO 客户端标识 CHAR(40) Y 通过建行生活APP下单场景，订单中客户端标识固定设为comccbpay1234567890cloudmerchant
TIMEOUT 订单超时时间 CHAR(14) N 格式：YYYYMMDDHHMMSS（如：20120214143005） 银行系统时间> TIMEOUT时拒绝交易，若送空值则不判断超时。  当该字段有值时参与MAC校验，否则不参与MAC校验。
PLATFORMPUB 服务方公钥 VARCHAR(256) Y 仅作为源串参加MD5摘要，不作为参数传递
MAC MD5加密串 CHAR(32) Y 采用标准MD5算法，对以上字段进行MAC加密（32位小写），由商户实现。
PLATFORMID 服务方编号 CHAR(16) Y 仅作为参数传递，不参与MAC校验
ENCPUB 商户公钥密文 VARCHAR(512) Y 使用服务方公钥对商户公钥后30位进行RSA加密后的密文，仅作为参数传递，不参与MAC校验
注：字符串中变量名必须是大写字母。
*/

type PayParams struct {
	MERCHANTID string `json:"MERCHANTID"` //商户代码 CHAR(15) Y 由建行统一分配
	POSID      string `json:"POSID"`      //柜台代码 CHAR(9) Y 由建行统一分配
	BRANCHID   string `json:"BRANCHID"`   //分行代码 CHAR(9) Y 由建行统一分配
	//POSID19    string `json:"POSID19"`    //商户19位终端号 CHAR(19) N 由建行统一分配，使用微信支付则必传，仅作为参数传递，不参与MAC校验
	ORDERID string `json:"ORDERID"` // 订单号 CHAR(30) Y 由商户提供，最长30位 PAYMENT 付款金额 NUMBER(16,2) Y 由商户提供，最长30位
	PAYMENT string `json:"PAYMENT"` //  付款金额 NUMBER(16,2) Y 由商户提供
	//CURCODE      string `json:"CURCODE"`      // 币种 CHAR(2) Y 缺省为01－人民币（只支持人民币支付）
	//TXCODE       string `json:"TXCODE"`       // 交易码 CHAR(6) Y 由建行统一分配为520100
	//REMARK1 string `json:"REMARK1"` // 备注1 CHAR(30) Y 网银不处理，直接传到城综网，该字段只支持送数字和英文
	//REMARK2 string `json:"REMARK2"` // 备注2 CHAR(30) Y 上送YS开头的服务方编号
	//TYPE     string `json:"TYPE"`     // 接口类型 CHAR(1) Y 1- 防钓鱼接口
	//GATEWAY  string `json:"GATEWAY"`  //网关类型 CHAR(100) Y 默认送0
	CLIENTIP string `json:"CLIENTIP"` //客户端IP CHAR(40) N 客户在商户系统中的IP
	REGINFO  string `json:"REGINFO"`  //客户注册信息 CHAR(256) N 客户在商户系统中注册的信息，中文需使用escape编码
	PROINFO  string `json:"PROINFO"`  //商品信息 CHAR(256) N 客户购买的商品，中文需使用escape编码
	//REFERER      string `json:"REFERER"`      //商户URL CHAR(100) N 商户送空值即可
	//THIRDAPPINFO string `json:"THIRDAPPINFO"` //客户端标识 CHAR(40) Y 通过建行生活APP下单场景，订单中客户端标识固定设为comccbpay1234567890cloudmerchant
	TIMEOUT     string `json:"TIMEOUT"`     //订单超时时间 CHAR(14) N 格式：YYYYMMDDHHMMSS（如：20120214143005） 银行系统时间> TIMEOUT时拒绝交易，若送空值则不判断超时。  当该字段有值时参与MAC校验，否则不参与MAC校验。
	PLATFORMPUB string `json:"PLATFORMPUB"` //服务方公钥 VARCHAR(256) Y 仅作为源串参加MD5摘要，不作为参数传递
	//MAC          string `json:"MAC"`          //MD5加密串 CHAR(32) Y 采用标准MD5算法，对以上字段进行MAC加密（32位小写），由商户实现。
	PLATFORMID string `json:"PLATFORMID"` //服务方编号 CHAR(16) Y 仅作为参数传递，不参与MAC校验
	//ENCPUB       string `json:"ENCPUB"`       //商户公钥密文 VARCHAR(512) Y 使用服务方公钥对商户公钥后30位进行RSA加密后的密文，仅作为参数传递，不参与MAC校验
}

// MacSign Md5签名 注意参数顺序，需要文档一致
func (p *PayParams) MacSign(qid string) (v string) {
	if p.TIMEOUT == "" {
		p.TIMEOUT = time.Now().Local().AddDate(0, 0, 1).Format("20060102150405")
	}

	s := fmt.Sprintf("MERCHANTID=%s&POSID=%s&BRANCHID=%s&ORDERID=%s&PAYMENT=%s&CURCODE=01&TXCODE=520100&REMARK1=&REMARK2=%s&TYPE=1&GATEWAY=0&CLIENTIP=172.0.0.1&REGINFO=%s&PROINFO=%s&REFERER=&THIRDAPPINFO=comccbpay1234567890cloudmerchant&TIMEOUT=%s&PLATFORMPUB=%s",
		p.MERCHANTID,
		p.POSID,
		p.BRANCHID,
		p.ORDERID,
		p.PAYMENT,
		p.PLATFORMID,
		p.REGINFO,
		p.PROINFO,
		p.TIMEOUT,
		p.PLATFORMPUB,
	)

	//logrus.WithField("qid", qid).Infof("MacSign 建行支付签名字符串: %s", s)

	v = fmt.Sprintf("%x", md5.Sum([]byte(s)))
	return
}

// Encpub ENCPUB : 各服务方使用自己的服务方公钥对商户公钥后30位进行RSA加密后，生成的密文串
func (p *PayParams) Encpub(mchPubKey string) (v string, err error) {
	v, err = RsaEncode(mchPubKey[len(mchPubKey)-30:], p.PLATFORMPUB)
	return
}

// PayInfo 支付参数
func (p *PayParams) PayInfo(conf *Config, qid string) (v string, err error) {
	if p.TIMEOUT == "" {
		p.TIMEOUT = time.Now().Local().AddDate(0, 0, 1).Format("20060102150405")
	}
	encpub, err := p.Encpub(conf.MchPubKey)
	if err != nil {
		return
	}
	v = fmt.Sprintf(
		"MERCHANTID=%s&POSID=%s&BRANCHID=%s&ORDERID=%s&PAYMENT=%s&CURCODE=01&TXCODE=520100&REMARK1=&REMARK2=%s&TYPE=1&GATEWAY=0&CLIENTIP=172.0.0.1&REGINFO=%s&PROINFO=%s&REFERER=&THIRDAPPINFO=comccbpay1234567890cloudmerchant&TIMEOUT=%s&MAC=%s&PLATFORMID=%s&ENCPUB=%s",
		p.MERCHANTID,
		p.POSID,
		p.BRANCHID,
		p.ORDERID,
		p.PAYMENT,
		p.PLATFORMID,
		p.REGINFO,
		p.PROINFO,
		p.TIMEOUT,
		p.MacSign(qid),
		p.PLATFORMID,
		encpub,
	)
	//logrus.WithField("qid", qid).Infof("PayInfo 建行支付下单参数: %s", v)
	return
}
