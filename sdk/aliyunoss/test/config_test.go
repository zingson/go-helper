package test

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"github.com/zingson/go-helper/sdk/aliyunoss"
	"io/ioutil"
)

func getConfig() (config *aliyunoss.Config) {
	// 读取阿里云AccessKey，从阿里云后台获取
	akjson, err := ioutil.ReadFile(".secret/access-key.json")
	if err != nil {
		log.Fatal(err.Error())
	}
	err = json.Unmarshal(akjson, &config)
	if err != nil {
		log.Fatal(err.Error())
	}
	return config
}
