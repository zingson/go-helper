package ccbh5_pay

import (
	"encoding/json"
	"errors"
)

// Query 服务方订单查询
func Query(conf *Config, params *QryParams) (cldBody *QueryCldBody, err error) {
	resBody, err := Post(conf, conf.ServiceOccplatreq, "svc_occPlatOrderQry", params)
	if err != nil {
		return
	}
	var resResult QryResResult
	err = json.Unmarshal([]byte(resBody), &resResult)
	if err != nil {
		return
	}
	if resResult.CLD_HEADER.CLD_TX_RESP.CLD_CODE != "CLD_SUCCESS" { // 判断成功状态码
		err = errors.New(resResult.CLD_HEADER.CLD_TX_RESP.CLD_DESC)
		return
	}
	cldBody = resResult.CLD_BODY
	return
}

type TxType string

const (
	TX_TYPE_0 TxType = "0" //支付，包括所有的支付/消费类功能
	TX_TYPE_1 TxType = "1" //退款，包括所有的退款/退货/撤销类功能
)

type QryParams struct {
	TX_TYPE             TxType `json:"TX_TYPE"`             //Y 交易类型， 0-支付，包括所有的支付/消费类功能 1-退款，包括所有的退款/退货/撤销类功能
	TXN_PRD_TPCD        string `json:"TXN_PRD_TPCD"`        //Y 查询时间范围类型， 06-近24小时内交易，99-自定义时间段查询 备注：06查询近24小时内交易，仅返回符合条件的最近一笔记录
	STDT_TM             string `json:"STDT_TM"`             //N 开始日期时间， 查询时间范围类型为99时必输，格式[yyyyMMddhhmiss]
	EDDT_TM             string `json:"EDDT_TM"`             //N 结束日期时间， 查询时间范围类型为99时必输，格式[yyyyMMddhhmiss]
	ONLN_PY_TXN_ORDR_ID string `json:"ONLN_PY_TXN_ORDR_ID"` //N 订单编号， 查询时间范围类型为06时必输
	SCN_IDR             string `json:"SCN_IDR"`             //N 场景标识 BHK - 本行卡；THK - 他行卡；ZFB - 聚合支付支付宝；CFT - 聚合支付微信
	PLAT_MCT_ID         string `json:"PLAT_MCT_ID"`         //N 服务商门店编号，外部平台商户号，不为空以这个字段为准
	CUSTOMERID          string `json:"CUSTOMERID"`          //N 商户号 ，建行商户编号，与外部平台商户号不能同时为空
	BRANCHID            string `json:"BRANCHID"`            //N 一级分行号，商户一级分行号，用建行商户编号时不能为空
	POS_CODE            string `json:"POS_CODE"`            //N 柜台号
	POS_ID              string `json:"POS_ID"`              //N POS终端编号
	TXN_STATUS          string `json:"TXN_STATUS"`          //Y 交易状态，00-交易成功标志；01-交易失败；02-不确定
	MSGRP_JRNL_NO       string `json:"MSGRP_JRNL_NO"`       //N 商户的流水号  商户支付流水号或者退款流水号；查询时间范围类型为06，且交易类型为1时必输
	PAGE                uint64 `json:"PAGE"`                //Y 当前页次,从1开始
}

type QueryCldBody struct {
	CUR_PAGE           string             `json:"CUR_PAGE"`           //当前页次，每页最多返回10条记录
	PAGE_COUNT         string             `json:"PAGE_COUNT"`         //总页次
	ED_CRD_PRTY_IDR_CD string             `json:"ED_CRD_PRTY_IDR_CD"` //商户号
	PY_AMT             string             `json:"PY_AMT"`             //支付金额
	MRCH_RFND_AMT      string             `json:"MRCH_RFND_AMT"`      //商户退款金额
	LIST               []QueryCldBodyList `json:"LIST"`
}

type QueryCldBodyList struct {
	ONLN_PY_TXN_ORDR_ID   string    `json:"ONLN_PY_TXN_ORDR_ID"`   //订单编号
	CLRG_STM_DT_TM        string    `json:"CLRG_STM_DT_TM"`        //交易时间 格式[yyyyMMddhhmiss]
	ACQ_FNDS_CLRG_DT      string    `json:"ACQ_FNDS_CLRG_DT"`      //记账日期
	ORDR_TM               string    `json:"ORDR_TM"`               //原支付订单时间
	AHN_TXNAMT            string    `json:"AHN_TXNAMT"`            //
	ORDR_PYRFD_AMT        string    `json:"ORDR_PYRFD_AMT"`        //
	TXN_CLRGAMT           string    `json:"TXN_CLRGAMT"`           //
	MRCHCMSN_AMT          string    `json:"MRCHCMSN_AMT"`          //
	ORIG_AMT              string    `json:"ORIG_AMT"`              //
	DISCOUNT_AMT          string    `json:"DISCOUNT_AMT"`          //
	RETGDS_ORIG_TXNAMT    string    `json:"RETGDS_ORIG_TXNAMT"`    //
	CST_ACCNO             string    `json:"CST_ACCNO"`             // 支付卡号
	CCYCD                 string    `json:"CCYCD"`                 // 币种
	TXN_STATUS            TxnStatus `json:"TXN_STATUS"`            // *状态 "00-交易成功标志  01-交易失败  02-不确定  04-不确定  TO-交易超时
	ORIOVRLSTTNEV_TRCK_NO string    `json:"ORIOVRLSTTNEV_TRCK_NO"` // 银行流水号
	MSGRP_JRNL_NO         string    `json:"MSGRP_JRNL_NO"`         // 退款时商户上送的退款流水号
	PAY_MODE              string    `json:"PAY_MODE"`              // BHK:建行;THK:他行;ZFB:支付宝;CFT:微信
	// 更多字段，参考建行文档
}

// 状态 "00-交易成功标志  01-交易失败  02-不确定  04-不确定  TO-交易超时
type TxnStatus string

const (
	TXN_STATUS_00 TxnStatus = "00"
	TXN_STATUS_01 TxnStatus = "01"
	TXN_STATUS_02 TxnStatus = "02"
	TXN_STATUS_T0 TxnStatus = "T0"
)

/*
CLD_HEADER
...CLD_TX_CHNL
...CLD_TX_TIME
...CLD_TX_CODE
...CLD_TX_SEQ
...CLD_TX_RESP
......CLD_CODE
......CLD_DESC
*/
type QryResResult struct {
	CLD_HEADER struct {
		CLD_TX_CHNL string `json:"CLD_TX_CHNL"`
		CLD_TX_TIME string `json:"CLD_TX_TIME"`
		CLD_TX_CODE string `json:"CLD_TX_CODE"`
		CLD_TX_SEQ  string `json:"CLD_TX_SEQ"`
		CLD_TX_RESP struct {
			CLD_CODE string `json:"CLD_CODE"` // 响应码
			CLD_DESC string `json:"CLD_DESC"`
		} `json:"CLD_TX_RESP"`
	} `json:"CLD_HEADER"`
	CLD_BODY *QueryCldBody `json:"CLD_BODY"`
}
