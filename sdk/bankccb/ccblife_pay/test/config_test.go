package test

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"github.com/zingson/go-helper/htime"
	"github.com/zingson/go-helper/sdk/bankccb/ccblife_pay"
	"os"
)

func init() {
	file, err := os.OpenFile(htime.NowF8()+".log", os.O_WRONLY|os.O_APPEND|os.O_CREATE|os.O_SYNC, 0600)
	if err != nil {
		panic(err)
	}
	logrus.SetOutput(file)
	logrus.SetFormatter(&logrus.TextFormatter{DisableQuote: true})
	logrus.SetLevel(logrus.DebugLevel)
}

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
