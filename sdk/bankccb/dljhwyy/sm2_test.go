package dljhwyy

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/tjfoc/gmsm/sm2"
	"github.com/tjfoc/gmsm/x509"
	"testing"
)

func TestGenKey(t *testing.T) {
	priKey, err := sm2.GenerateKey(rand.Reader) // 生成密钥对
	if err != nil {
		fmt.Println(err)
		return
	}
	pubKey := &priKey.PublicKey

	priKeyByte, err := x509.MarshalSm2UnecryptedPrivateKey(priKey)
	if err != nil {
		return
	}
	pubKeyByte, err := x509.MarshalSm2PublicKey(pubKey)
	if err != nil {
		return
	}
	t.Logf("SM2私钥：%s", base64.StdEncoding.EncodeToString(priKeyByte))
	t.Logf("SM2公钥：%s", base64.StdEncoding.EncodeToString(pubKeyByte))
}

func TestSm2Decode(t *testing.T) {
	prib64 := "MIGTAgEAMBMGByqGSM49AgEGCCqBHM9VAYItBHkwdwIBAQQg5qUWfiMsx/w9Jl1LOwehL+5H/Tuqjg9SHuZRt8+uZuygCgYIKoEcz1UBgi2hRANCAATI5dWSwLgWc3nuek7jgsJgecGQfqTlk9wth/DoF5Di6GpQoF0Db/wB2+XHHN8/L4Gq1/j+9JOCDxJsIRftgl0h"
	pubb64 := "MFkwEwYHKoZIzj0CAQYIKoEcz1UBgi0DQgAEyOXVksC4FnN57npO44LCYHnBkH6k5ZPcLYfw6BeQ4uhqUKBdA2/8AdvlxxzfPy+Bqtf4/vSTgg8SbCEX7YJdIQ=="
	data := "CCB"

	v, err := Sm2Encode(data, pubb64)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("加密结果：%s", v)

	v, err = Sm2Decode(v, prib64)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("解密结果：%s", v)
}

func TestSm2Encode(t *testing.T) {
	v, err := Sm2Encode("", "")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(v)
}
