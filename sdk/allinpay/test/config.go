package test

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"root/src/sdk/allinpay"
)

func config() (cfg *allinpay.Config) {
	b, err := ioutil.ReadFile(".pwd/allinpay-55233204816VEVU.json")
	if err != nil {
		log.Error(err)
		panic(err)
	}
	err = json.Unmarshal(b, &cfg)
	if err != nil {
		panic(err)
	}
	return
}
