package test

import (
	_ "embed"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"github.com/zingson/go-helper/sdk/metro"
)

//go:embed .secret/wx-production.json
var configStr string // 生产环境

////go:embed .secret/test.json
//var configStr string // 测试环境

var config *metro.Config

func init() {
	err := json.Unmarshal([]byte(configStr), &config)
	if err != nil {
		logrus.Error(err.Error())
	}
}
