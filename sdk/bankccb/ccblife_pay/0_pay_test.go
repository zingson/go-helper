package ccblife_pay

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"testing"
)

func getConfig() (conf *Config) {
	_, err := toml.DecodeFile("D:\\Projects\\hlib-go\\helper\\sdk-bankccb\\ccblife_pay\\.secret\\config.toml", &conf)
	if err != nil {
		fmt.Println("error:", err.Error())
		return
	}
	return
}

func TestSvcOccPlatOrderQry(t *testing.T) {
	conf := getConfig()
	_, err := Query(conf, &QryParams{
		TX_TYPE:             TX_TYPE_0,
		TXN_PRD_TPCD:        "99",
		STDT_TM:             "20220415000000",
		EDDT_TM:             "20220417000000",
		ONLN_PY_TXN_ORDR_ID: "",
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
}

func TestSvcOccPlatOrderQry2(t *testing.T) {
	conf := getConfig()
	cldBody, err := Query(conf, &QryParams{
		TX_TYPE:             TX_TYPE_0,
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
	cldBody, err := Refund(conf, "15155562905736929293", 1, "20220417142653")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(cldBody.ORDER_NUM)
	t.Log("refund success")
}
