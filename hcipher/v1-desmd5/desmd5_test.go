package desmd5

import (
	"net/url"
	"testing"
)

func TestEncryptDesMd5(t *testing.T) {
	key := "88888888"
	s := `{"appid":"88888888"}`
	/*
		密文 e2a19818214e2c1523208747dd872fa9xmF4oTvCx2mrAp8erL7oHg==
		明文 123中文
	*/

	//加密
	v, e := EnDesMd5(s, key)
	if e != nil {
		t.Error(e.Error())
		return
	}
	t.Log(v)
	t.Log(url.QueryEscape(v))

	// 解密
	v, e = DeDesMd5(v, key)
	if e != nil {
		t.Error(e.Error())
		return
	}
	t.Log(v)
}
