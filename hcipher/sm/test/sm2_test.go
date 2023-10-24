package test

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/tjfoc/gmsm/sm2"
	"github.com/tjfoc/gmsm/x509"
	"testing"
)

// 同 java的签名验签 验证通过
func TestHashHex3(t *testing.T) {

	const pubKey = "MFkwEwYHKoZIzj0CAQYIKoEcz1UBgi0DQgAEoo5iYxcb5VM1GSHL5drRr9KsWTBS4IMYzO4UReQikWIRWhXrOk8y6WqlGtR+XoQ61sZOM+YG4XEr2jPQTOQ7gg=="
	const priKey = "MD0CAQAwCwYHKoZIzj0CAQUABCswKQIBAQQgoGTIw/VXHOrRQI+BjsT04H7v6JZsDCcA2bmu9i4FcHGgAgUA"

	sValue := "busines&certId=BANKTEST001&orderId=1669271981244&signType=SM2&txnTime=20221124063941&txnType=H5&version=1.0.1"

	hexD, err := base64.StdEncoding.DecodeString(priKey)
	if err != nil {
		t.Error(err)
		return
	}
	privateKey, err := x509.ParsePKCS8UnecryptedPrivateKey(hexD)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log("SM2签名：")
	signBytes, err := privateKey.Sign(rand.Reader, []byte(sValue), nil)
	if err != nil {
		t.Error(err)
		return
	}
	signResult := base64.StdEncoding.EncodeToString(signBytes)
	t.Log("sign:" + signResult)

	t.Log("SM2验签：")
	hexPub, err := base64.StdEncoding.DecodeString(pubKey)
	if err != nil {
		t.Error(err)
		return
	}
	publickKey, err := x509.ParseSm2PublicKey(hexPub)
	if err != nil {
		t.Error(err)
		return
	}
	signResult = "MEQCIB8ecIC8z8unHF0e4/Ecn/erneFlWmWVA5p0RaNFukJ2AiBpLSPESPDHRkyNRYfQTxngmIbNVHFLu41ZZLJmjCvQfw=="
	signBytes2, err := base64.StdEncoding.DecodeString(signResult)
	if err != nil {
		t.Error(err)
		return
	}
	bl := publickKey.Verify([]byte(sValue), signBytes2)
	t.Log(bl)
}

// 生成密钥对
func TestGenKey(t *testing.T) {
	priKey, err := sm2.GenerateKey(rand.Reader)
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
