package v2

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

// OrderQueryV2 3.4订单查询接口   http://*:*/Order/QueryV2（请从供货商处获取）
func (c *Client) OrderQueryV2(mOrderId string) (result *OrderQueryV2Result, err error) {
	timestamp := Timestamp()
	sign := Md5Sign(c.Cfg.AppKey + timestamp + mOrderId + c.Cfg.AppSecret)
	err = c.Request("/Order/QueryV2",
		fmt.Sprintf("AppKey=%s&TimesTamp=%s&Sign=%s&MOrderID=%s&OrderID=", c.Cfg.AppKey, timestamp, sign, mOrderId), &result)
	if err != nil {
		return
	}
	if result.Code != ERR_SUCCESS.Code {
		err = Err(result.Code, result.Msg)
		return
	}
	if result.Sign != Md5Sign(c.Cfg.AppKey+strconv.FormatInt(result.TimesTamp, 10)+strconv.FormatInt(result.Code, 10)+strconv.FormatInt(int64(result.Data.OrderState), 10)+c.Cfg.AppSecret) {
		err = ERR_SIGN
		return
	}
	err = extendParamRsaDecrypt(result, c.Cfg.RsaPriKey)
	if err != nil {
		log.WithField("orderId", mOrderId).Error(err.Error())
		err = ERR_RS_DECRYPTY
		return
	}
	return
}

// 卡密提取订单时解密卡密
func extendParamRsaDecrypt(result *OrderQueryV2Result, priKey string) (err error) {
	if result == nil || result.Data == nil || result.Data.ExtendParam == nil {
		return
	}
	if result.Data.ExtendParam.ChannelSerialNumber != "" {
		v, err := RsaDecrypt(result.Data.ExtendParam.ChannelSerialNumber, priKey)
		if err != nil {
			return err
		}
		result.Data.ExtendParam.ChannelSerialNumber = strings.TrimSpace(v)
	}
	if result.Data.ExtendParam.CardPwd != "" {
		v, err := RsaDecrypt(result.Data.ExtendParam.CardPwd, priKey)
		if err != nil {
			return err
		}
		result.Data.ExtendParam.CardPwd = strings.TrimSpace(v)
	}
	if result.Data.ExtendParam.CardNumber != "" {
		v, err := RsaDecrypt(result.Data.ExtendParam.CardNumber, priKey)
		if err != nil {
			return err
		}
		result.Data.ExtendParam.CardNumber = strings.TrimSpace(v)
	}
	return
}

type OrderQueryV2Result struct {
	Code      int64
	Msg       string
	TimesTamp int64
	Sign      string
	Data      *OrderQueryV2ResultData
}

type OrderQueryV2ResultData struct {
	OrderID        int64
	MOrderID       string
	OrderState     int64
	ChargeAccount  string
	BuyCount       int64
	Price          int64
	SellDebitAmout int64
	SellRebate     int64
	CreateTime     string
	ExtendParam    *OrderQueryV2ResultDataExtendParam
}

type OrderQueryV2ResultDataExtendParam struct {
	CardDeadline        string
	CardNumber          string
	CardPwd             string
	ChannelSerialNumber string
	FinishTime          string
	OfficialDes         string
	OfficialOrderID     string
}
