package v3_test

import (
	_ "embed"
	"fmt"
	"github.com/BurntSushi/toml"
	v3 "root/src/sdk/wxpay/v3"
	"testing"
)

//go:embed .secret/1579635811.toml
var payToml string

func Client() *v3.Client {
	var option v3.Option
	_, err := toml.Decode(payToml, &option)
	if err != nil {
		panic(err)
	}
	c, err := v3.NewClient(option)
	if err != nil {
		panic(err)
	}
	fmt.Println(c.Mchid)
	return c
}

func TestWxPubKey(t *testing.T) {
	c := Client()
	c.GetWxPubKey("")
}
