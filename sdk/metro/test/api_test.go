package test

import (
	_ "embed"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"github.com/zingson/go-helper/sdk/metro"
	"testing"
)

// 根据手机号开票测试
func TestTicketOpen(t *testing.T) {
	mobile := "13611703040"
	productCode := "测试"

	// 手机产品取码
	tc, err := metro.TicketOpen(config, mobile, productCode, metro.Rand32())
	if err != nil {
		t.Error(err)
		return
	}
	logrus.Info("计次票信息：", tc.JSON())

	// 生成乘车链接
	url, err := metro.Entry(config, tc.TicketCode, mobile)
	if err != nil {
		t.Error(err)
		return
	}
	logrus.Info("乘车码链接 " + url)
	t.Log("success...")
}

// 授权
func TestAuthByMobile(t *testing.T) {
	userId, err := metro.AuthByMobile(config, "13611703040")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("userId:", userId)
	t.Log("success...")
	// 手机号授权测试结果：OK 正常返回用户ID b89b4187202240b7a49007901305a17b
}

// 查产品信息
func TestProductInfo(t *testing.T) {
	prod, err := metro.ProductInfo(config, "测试")
	if err != nil {
		t.Error(err)
		return
	}
	pbytes, _ := json.Marshal(prod)
	t.Log(string(pbytes))
	t.Log("success...")
	// 产品信息查询测试结果：返回产品信息比接口文档定义的多,且字段名与文档不一致，如：可乘车次数文档字段名 times,实际返回没有
}

// 购票
func TestMonthlyTicketOpen(t *testing.T) {
	_, err := metro.MonthlyTicketOpen(config, "b89b4187202240b7a49007901305a17b", "d001", metro.Rand32())
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("success...")
}

// 计次票二维码H5页面嵌入
func TestEntry1(t *testing.T) {
	url, err := metro.Entry(config, "1671344181295058944", "13611703040")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(url)
	// 测试结果：
}

func TestEntry2(t *testing.T) {
	url, err := metro.Entry(config, "1347006618226790400", "13611703040")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(url)
	// 测试结果：
}
