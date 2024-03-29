package example

import (
	"github.com/zingson/go-helper/sdk/uppay/ocapp"
	"testing"
)

// 获取银联平台公钥
func TestPubkey(t *testing.T) {
	_, err := ocapp.Pubkey(cfg)
	if err != nil {
		t.Error(err.Error())
	}
	t.Log("success................")
}

// 获取银联平台公钥2
func TestPubkey2(t *testing.T) {
	_, err := ocapp.Pubkey(cfg821330248164060)
	if err != nil {
		t.Error(err.Error())
	}
	t.Log("success................")
}
