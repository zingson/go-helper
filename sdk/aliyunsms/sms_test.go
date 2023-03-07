package aliyunsms

import (
	"github.com/BurntSushi/toml"
	"testing"
)

func TestAliyunSmsConfig_SendSms(t *testing.T) {
	var config *Config
	_, err := toml.DecodeFile("D:\\Projects\\hlib-go\\helper\\sdk-aliyunsms\\.secret\\广鲲数藏.toml", &config)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(config)
	err = Send(config, "", "13611703040", "广鲲数藏", "SMS_238300675", map[string]string{"code": "1234"})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("Success...")
}
