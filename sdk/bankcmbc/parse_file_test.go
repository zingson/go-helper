package bankcmbc

import (
	_ "embed"
	"encoding/json"
	"testing"
)

//go:embed YHK_MS_20210929.txt
var yhkms string

var rsapub = `
-----BEGIN RSA PUBLIC KEY-----
MIGJAoGBAIlY/pKhxKAjym0zjem4hq2U3Eh9ToCrfKdDReTPe0EAMjeW3bQIZJLD
KuITynWTrZJ95+2I01ppiIn+VedYi/uJegI7/F1romDcv3NgEpsoZntEinyymOMS
1Azlp+u5C/1j1SVaavbULkmF39dsaXBiEifYTKGT0U3EEfDCPni5AgMBAAE=
-----END RSA PUBLIC KEY-----
`
var rsapri = `
-----BEGIN RSA PRIVATE KEY-----
MIICYAIBAAKBgQCJWP6SocSgI8ptM43puIatlNxIfU6Aq3ynQ0Xkz3tBADI3lt20
CGSSwyriE8p1k62SfeftiNNaaYiJ/lXnWIv7iXoCO/xda6Jg3L9zYBKbKGZ7RIp8
spjjEtQM5afruQv9Y9UlWmr21C5Jhd/XbGlwYhIn2Eyhk9FNxBHwwj54uQIDAQAB
AoGAamTM9ytm1CJFeagZA3bUpPwOU/z1Zcjxi+QZ7XAn6ydKvzMX1JE3z1RuEKkC
CWh3aWYs1h1Kk9vyT+r7Pe2MWQ9GLEpDw/nP9XNfmFeUdk5m+m8im9YOf5/PrFZM
YZlBOWUUuqngbtfakxxeMyUDlP4QBFVykBoteCObrLqsp10CRQD9tL7CSi4dMrSL
O4NtNJ9AxeWPHun2T/MFgcbaCd2ERHRKlWZsGu8k3sGG/8m/93/ZH1OpsjWQghWY
GtzJD3HzsFUACwI9AIqW6e7CBTKZht0WG5+U6jLUY/yKls+ndzKm1Y/w/h2Jr3MI
e/Z4v9JT5iF4UOImtPoWvGcXOrW4fPxQywJEUELm+lY3YntRDJ8mQ90a6IXyyqVQ
BOFkE4Dr5LysPJTfaVz8SwT2VOa3uLqhG77zzj+P2yaKtY3BwR32bREazqohKeMC
PGQdduqKcFTIQXuOz++tFK4Zbg1uVFm34UzO5nHwJrJR11OjKmG3guK+xv0gvFVS
nuQW7o0OY9QIbQmclwJFAIlPzvwU7t4pUYt7P0ApaenWU+p6HR+zdBJyVEvfeW+L
59avveUOjCajGOHIoihinsNsXl9aJ8aTZ6QUqMOnWj0+vcaD
-----END RSA PRIVATE KEY-----
`

// 解析文件测试
func TestParseFile(t *testing.T) {
	rows, err := ParseFile(yhkms, rsapub)
	if err != nil {
		t.Fatal(err)
	}
	rbytes, _ := json.Marshal(rows)
	println("解析结果转为JSON格式内容：")
	println(string(rbytes))
}
