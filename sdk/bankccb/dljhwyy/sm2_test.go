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
	/*
		Java demo
		生成的公钥：MFkwEwYHKoZIzj0CAQYIKoEcz1UBgi0DQgAEGdwOnkpfJqP8DMxuaNnB7U2YP2jWZ6UK/nGcX0E23lsapo5e/mgBQ8thVhOpIgDZ8BdYmF0+Xf3Et9RDQuv2YA==
		生成的私钥：MIGTAgEAMBMGByqGSM49AgEGCCqBHM9VAYItBHkwdwIBAQQgWmuKc/BhVTK4ztyzkETcI0qFxwvZGnUYelCrydd18kqgCgYIKoEcz1UBgi2hRANCAAQZ3A6eSl8mo/wMzG5o2cHtTZg/aNZnpQr+cZxfQTbeWxqmjl7+aAFDy2FWE6kiANnwF1iYXT5d/cS31ENC6/Zg
		加密字符串：BLy1ToTgaztxWsteCwDfGxoLaOo94imjdcjBGWm6lajnGvKKfr6KwKkMndDECeVateW1A2A740J1QiVuVSAU1jhFJHvXugUQ/Oni7scJzVAW5e6IgTYntEkO+s4DMwkyDWD1Uw==
		解密字符串：CCB
	*/

	prib64 := "MIGTAgEAMBMGByqGSM49AgEGCCqBHM9VAYItBHkwdwIBAQQgWmuKc/BhVTK4ztyzkETcI0qFxwvZGnUYelCrydd18kqgCgYIKoEcz1UBgi2hRANCAAQZ3A6eSl8mo/wMzG5o2cHtTZg/aNZnpQr+cZxfQTbeWxqmjl7+aAFDy2FWE6kiANnwF1iYXT5d/cS31ENC6/Zg"
	pubb64 := "MFkwEwYHKoZIzj0CAQYIKoEcz1UBgi0DQgAEGdwOnkpfJqP8DMxuaNnB7U2YP2jWZ6UK/nGcX0E23lsapo5e/mgBQ8thVhOpIgDZ8BdYmF0+Xf3Et9RDQuv2YA=="
	data := "CCB"

	v, err := Sm2Encode(data, pubb64)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("加密结果：%s", v)

	v = "BLy1ToTgaztxWsteCwDfGxoLaOo94imjdcjBGWm6lajnGvKKfr6KwKkMndDECeVateW1A2A740J1QiVuVSAU1jhFJHvXugUQ/Oni7scJzVAW5e6IgTYntEkO+s4DMwkyDWD1Uw=="
	r, err := Sm2Decode(v, prib64)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("解密结果：%s", r)
}

func TestSm2Encode(t *testing.T) {
	v, err := Sm2Encode("", "")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(v)
}
