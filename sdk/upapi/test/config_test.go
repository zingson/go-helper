package test

import (
	_ "embed"
	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
	"github.com/zingson/go-helper/sdk/upapi"
	"os"
	"testing"
)

func init() {
	file, err := os.OpenFile("t.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE|os.O_SYNC, 0600)
	if err != nil {
		panic(err)
	}
	logrus.SetOutput(file)
	logrus.SetFormatter(&logrus.TextFormatter{DisableQuote: true})
	logrus.SetLevel(logrus.DebugLevel)
}

// 注意：toml字段名需要与Config结构体属性名大小写一致. .secret文件夹未提交到git，需要自行创建

//go:embed .secret/宁波银联-闪券发券.toml
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
