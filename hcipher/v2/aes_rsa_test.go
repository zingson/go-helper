package v2

import "testing"

func TestEnAesRsa(t *testing.T) {
	pubKey := "MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCf0arfpOm8uRJ/HUg3HUzGnvXF9L18tlFvEAas59XZ5GAxQSIVyK3+yAxo4YQ1zjcHCFDk0zvEHvb7MPNSulu6zxDoyskCLVhTfzmWncewyAtu1YZGpIZaI496JtBsDjb2uszR8FrPhDoC4tyiEDVGMyGDa2hruQfXExrMSymY/wIDAQAB"
	plaintext := "123"

	v, err := EnAesRsa(plaintext, pubKey)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("加密结果：" + v)

	priKey := "MIICdgIBADANBgkqhkiG9w0BAQEFAASCAmAwggJcAgEAAoGBAJ/Rqt+k6by5En8dSDcdTMae9cX0vXy2UW8QBqzn1dnkYDFBIhXIrf7IDGjhhDXONwcIUOTTO8Qe9vsw81K6W7rPEOjKyQItWFN/OZadx7DIC27Vhkakhlojj3om0GwONva6zNHwWs+EOgLi3KIQNUYzIYNraGu5B9cTGsxLKZj/AgMBAAECgYBYyOnchIk1Rrt30pSjyas1TGnNJ0F9Xuuuf4C13bV37t9hfYWqBGpk+E2sllwUaBM53OzTsmWpdmpO8cq7UrxSEqkOJU4zUtdD7031eTgq9Fyp53EBU54mKOAOVoBoHdu/Qfv6GI8X2vmlf6FFeqheWST/AZ/XPkbG5Pp3zAa5QQJBAMpm7usIR3Vdqyl1OVBm4IU+eFTdDDPl31T/CR7jaUAH+d9KFedJOM2JdyQPwKFsRTwEEwBm9vgNrk3Kq6PAEQsCQQDKI/rJG/b04ZP7af28qPRmEquozdevsDe0s33dzWhLimqeriHDI8+uhcqgGIC1gh62LWnYpdEE8VkCDg4LmzhdAkEArMTskEeC59ZK8pqTj+QhJtvKT3ZYojxIRP9mQ62O/A9S5Z0R0VmZWSlMP8YKgkAvYSmBJsi9a8QR02l61c5vPQJAbr3HZuYrJX1v1Qz8NZ9aRZF0+cXLpDSmUBkFm74spTXvs38yf/XekX46w/qoiMgAi03V7xroqAyQ9s88Yp9nAQJAXB3jqkPjSpeDzn09JqjdKJ5XzqSidA4/vhZotbIC24rDBN1+dcB67Lx2ggicZYrjpc70CjBgoBR+I6w7qpatQg=="
	r, err := DeAesRsa(v, priKey)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("解密结果：" + r)
}
