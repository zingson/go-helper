package test

import (
	"encoding/json"
	"github.com/zingson/go-helper/sdk/bankccb/ccblife_pay"
	"os"
	"testing"
)

func getConfig() (conf *ccblife_pay.Config) {
	b, err := os.ReadFile("./.secret/config.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(b, &conf)
	if err != nil {
		panic(err)
	}
	return
}

func TestSvcOccPlatOrderQry2(t *testing.T) {
	conf := getConfig()
	cldBody, err := ccblife_pay.Query(conf, &ccblife_pay.QryParams{
		TX_TYPE:             ccblife_pay.TX_TYPE_0,
		TXN_PRD_TPCD:        "06",
		STDT_TM:             "",
		EDDT_TM:             "",
		ONLN_PY_TXN_ORDR_ID: "15155628417745510412",
		SCN_IDR:             "",
		PLAT_MCT_ID:         "",
		CUSTOMERID:          conf.MerchantId,
		BRANCHID:            conf.BranchId,
		POS_CODE:            "",
		POS_ID:              "",
		TXN_STATUS:          "00",
		MSGRP_JRNL_NO:       "",
		PAGE:                1,
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(cldBody.LIST[0].TXN_STATUS)

}

// 退款测试
func TestRefund(t *testing.T) {
	conf := getConfig()
	cldBody, err := ccblife_pay.Refund(conf, "15155562905736929293", 1, "20220417142653")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(cldBody.ORDER_NUM)
	t.Log("refund success")
}

// 测试订单推送
func TestSVC_occMebOrderPush(t *testing.T) {
	conf := getConfig()
	ccblife_pay.SVC_occMebOrderPush(conf, &ccblife_pay.OrderPushParams{
		USER_ID:          "",
		ORDER_ID:         "",
		ORDER_DT:         "",
		TOTAL_AMT:        "",
		PAY_AMT:          "",
		DISCOUNT_AMT:     "",
		ORDER_STATUS:     "",
		REFUND_STATUS:    "",
		MCT_NM:           "",
		CUS_ORDER_URL:    "",
		OCC_MCT_LOGO_URL: "",
	})
}
