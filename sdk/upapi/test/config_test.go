package test

import (
	_ "embed"
	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
	"github.com/zingson/go-helper/sdk/upapi"
	"os"
	"testing"
	"time"
)

func init() {
	_ = os.MkdirAll("logs/"+time.Now().Format("200601"), 0600)
	file, err := os.OpenFile("logs/"+time.Now().Format("200601")+"/"+time.Now().Format("20060102T15")+".log", os.O_WRONLY|os.O_APPEND|os.O_CREATE|os.O_SYNC, 0600)
	if err != nil {
		panic(err)
	}
	logrus.SetOutput(file)
	logrus.SetFormatter(&logrus.TextFormatter{DisableQuote: true})
	logrus.SetLevel(logrus.DebugLevel)
}

// 注意：toml字段名需要与Config结构体属性名大小写一致. .secret文件夹未提交到git，需要自行创建

//go:embed .secret/大连银联-闪券发券.toml
var config string

// cfgtoml 通过开发环境配置接口参数，测试验证接口
func cfgtoml() (c *upapi.Config) {
	_, err := toml.Decode(config, &c)
	if err != nil {
		panic(err)
	}
	return
}

func TestConfig(t *testing.T) {
	t.Log(config)
}
