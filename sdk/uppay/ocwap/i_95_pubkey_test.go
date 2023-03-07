package ocwap

import (
	"testing"
)

// 获取银联平台公钥
func TestPubkey(t *testing.T) {
	_, err := Pubkey(_cfg())
	if err != nil {
		t.Error(err.Error())
	}
	t.Log("success................")
}
