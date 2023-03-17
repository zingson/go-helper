package test

import (
	"github.com/zingson/go-helper/sdk/bankccb/ccblife_pay"
	"testing"
)

// TestSvcOccPlatOrderQry 订单查询测试
func TestSvcOccPlatOrderQry(t *testing.T) {
	conf := getConfig()
	cldBody, err := ccblife_pay.PlatOrderQry(conf, &ccblife_pay.QryParams{
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
