package aliyunoss

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
)

func getConfig() (config *Config) {
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
