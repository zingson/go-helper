package aliyunsls

import (
	_ "embed"
	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
	"testing"
	"time"
)

//go:embed .secret/sls.toml
var slsToml string

func TClient() *Client {
	var v Option
	_, err := toml.Decode(slsToml, &v)
	if err != nil {
		panic(err)
	}
	return NewClient(v)
}

func TestLocalIP(t *testing.T) {
	t.Log(LocalIP())
}

func TestLog(t *testing.T) {
	TClient().Log("helper", map[string]string{"msg": "helper."})
	time.Sleep(time.Minute)
}

func TestHook(t *testing.T) {
	AddHook("helper", TClient())
	time.Sleep(3 * time.Second)
	logrus.Info("测试 logrus hook")
	time.Sleep(time.Minute)
}
