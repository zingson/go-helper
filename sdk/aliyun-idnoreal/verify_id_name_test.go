package aliyun_idnoreal

import (
	_ "embed"
	"testing"
)

// //go:embed .appcode
var appcode string

var conf = &Config{
	ServiceUrl: "https://dfidveri.market.alicloudapi.com",
	AppCode:    appcode,
}

// 实名认证测试
func TestCertification(t *testing.T) {
	err := Certification(conf, "421124199306257017", "必白")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("success....")
}
