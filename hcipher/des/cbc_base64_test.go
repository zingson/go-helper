package des

import "testing"

func TestEncodeCBCBase64(t *testing.T) {
	kye := "12345678"
	v, err := EncodeCBCBase64("132", kye)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log("加密结果：", v)

	v, err = DecodeCBCBase64(v, kye)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("解密结果：", v)
}
