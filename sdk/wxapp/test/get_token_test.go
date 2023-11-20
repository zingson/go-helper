package test

import (
	_ "embed"
	"github.com/BurntSushi/toml"
	"github.com/zingson/go-helper/sdk/wxapp"
	"testing"
)

//go:embed .secret/wx540a417fdf273755.toml
var config string

func cfg() (c *wxapp.Config) {
	_, err := toml.Decode(config, &c)
	if err != nil {
		panic(err)
	}
	return
}

var accessToken = "74_sZQZwnMbzhvtHjnMBK99Dbb3uRJN0d04jt4NS_8XVW6hY6HTIG-BTZQS0KuCTi5B2wtj1V0PX1P0VMFAhrF9OMWP45XNzriXCM3ag3wAvO1ozEMdeigmFtBPCAQDMXbAGAOXI"

func TestGetToken(t *testing.T) {
	wxapp.Token(cfg())
}
