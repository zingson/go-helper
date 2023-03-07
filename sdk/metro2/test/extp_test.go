package test

import (
	"github.com/sirupsen/logrus"
	"os"
	"root/src/sdk/metro2"
	"testing"
)

var config = &metro2.Config{
	Appid:       "SM202208080004",
	ServiceUrl:  "https://extptp.test.smartmetro.tech",
	AesSecret:   "8689712316164033",
	AesIv:       "94817a2360c4e8b2",
	MetroRsaPub: "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAwQ1hK/gT39XhFok2tWPs6yugP3p9wOkughpcrg3OpwhvklBxQbk7Bq6YASMa/jWqgeb/xCDx1jEMtXbMpEClWjaVsqLQ4VLVO+Lj0m4MTXJBdGz2yTtjr00CSLtMEgpSYRbsv2iWTCvROqErr3bakGt/zG3kmp/6cg5N/pm2CvBtrFK8kUPBmZHjz9uSC5LQ6eL9Hrb7E2qXSA0DEFcmCyptYMqWxun5mNXo7+Ijc0dDCL+ewo4wcHSVP8bSmlOamhQiiFUPAJXWRpaPayiCiJ0G0oHyNbp2ctwQyjGQkpslO+T3PKjeXgfv/daQg88yxX4M3dRNjUf0M9UHSEqoBQIDAQAB",
	H5url:       "https://extph5.test.smartmetro.tech/mini/mini-ticket-detail.html",
	RsaPri: `-----BEGIN PRIVATE KEY-----
MIIEvwIBADANBgkqhkiG9w0BAQEFAASCBKkwggSlAgEAAoIBAQDVzh8Ck5ThNpKd
dpnOCcAncjUwg/51chchv8zEALx7j6jMB25WxMtkcEhvAn680OCB1YE71562vS+R
umdRLsI5xT6fQRanO5isstscje71weqnvew8eOtwZ/xGWrBW4UwWY8YzI7pyCton
BR3Hu+BCcIyM/YoesxGD4Af11G91vi/K52h+6GFucoUwZXI4GXvXmhvxYOi8kzGr
GNS7CYf7sWmDT6Nie0OSnzzbjZXEA5hJqG2TbpDCle7m2r1a0x9J+5GolFetiXk1
qwsW1CJfCga8B3U6Og/bLrtZbMSMLcFSqgzLM/MBsm4hskuvCKXGHKWielOJJ+FJ
43doR8xPAgMBAAECggEATsDZ000hLcO4RaAGD0qwKNd8mB5GgGRB+QB5IElVI/5K
oryp+/QJJznktF8q58nYqHpIXA6UO6N7/iW3IMQkbrsk0exbt8XP+uz2oJH3Tzba
hGcEuVGhB4qF9jQ5eAcMy8J8oyGp74/nwy+zRHsDo/VpEBXj5mL3NKg0xmw1khyL
GEUhnt8GA8dumLKwAuG2/t9zfFsNdZTjIQxbHOHWc0CqedospbbWgXtlUBaqqjyp
vxkKXvlErFCkPIyK7wOcmQtyyF8m3IaY/NfeM0TVv2f0lHwVpUXTTtOIymA1o6Ro
YAsN1gtPfJMXfmA9OPzhelbbNnAp815qOALRS3INYQKBgQD7hnF7bL49e67DPHrs
AkatXG7hp4TmLj6Gn5CfemUDzkovvLIH8VR5feZAcS+Z4jWBagvJlN4siUxLUpI2
9W19w3y3Mbd3EE8YB9pNFx0Uz0/ADIVPnSd5w0Ok1B4sRu9d68UM7RxYG+SO/Nj4
ZKTZ3HmnSvF9R8kROxqufsMXIwKBgQDZm+JfOJbvs1UlQTuUckuDAteNThMNBtx8
qQpDfgNfbkWGyE423gE0ICfXF5ClS3p7Mo3n0pRBYXy0tDd6Wo2YDxW9iUtuFfAL
il109k7nK2Nb+bFLZQrcHuMo3jUEDy4OhCt69NvyiXj9dUdSmU9IKrRIu0TK2XtY
EgojlJoe5QKBgQCFbMsUKxo7qTmKrbGTMp3lZqwXHfMc25klds2UG6wsOakW66UR
G46xJ/0VYDVdDydM8EEyfLriqy6ColmXt9eOKD4nO8NT8J3UZI4D6OfvWw22Fa4+
DmKbb96ZOECNQk/F5cTQ15L6lklHJI/ALDtaql8KRHIYABWEA3Ni7zF0OwKBgQCT
YDiq5qeKhXj39zsDqXisrOMRRzwtyWTQZKeX2CMuoX8i7kvSav6Dr/drfAExgXHd
N/rVc0+HDCAqPheInQViY20E5ZQZZXAiUL5EtX/wnfj31J6XgkIdnCmahwt+yU0W
9bqA1o6Trzkq5x+7uCrypEFfNL09aJdZqTYGrODIZQKBgQDDRIYu8dQC9Qkp8BNx
h30JvR5152k1B+J0V12rSpupkO1eiHiWYH86BVBEJQN9udzm05LjF9A8WUS9818h
MxxNy4nBisEp6hZk4pIGwg2cBSAGanzOc1Sq3rcWrpvpwtEbzqWR+Ha4I+ZtmEyM
dnwS7DiKJREMlCJ95/n9eTq8pQ==
-----END PRIVATE KEY-----
`,
	RsaPub: `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA1c4fApOU4TaSnXaZzgnA
J3I1MIP+dXIXIb/MxAC8e4+ozAduVsTLZHBIbwJ+vNDggdWBO9eetr0vkbpnUS7C
OcU+n0EWpzuYrLLbHI3u9cHqp73sPHjrcGf8RlqwVuFMFmPGMyO6cgraJwUdx7vg
QnCMjP2KHrMRg+AH9dRvdb4vyudofuhhbnKFMGVyOBl715ob8WDovJMxqxjUuwmH
+7Fpg0+jYntDkp88242VxAOYSahtk26QwpXu5tq9WtMfSfuRqJRXrYl5NasLFtQi
XwoGvAd1OjoP2y67WWzEjC3BUqoMyzPzAbJuIbJLrwilxhylonpTiSfhSeN3aEfM
TwIDAQAB
-----END PUBLIC KEY-----
`,
}

func init() {
	file, err := os.OpenFile("t.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE|os.O_SYNC, 0600)
	if err != nil {
		panic(err)
	}
	logrus.SetOutput(file)
	logrus.SetFormatter(&logrus.TextFormatter{DisableQuote: true})
	logrus.SetLevel(logrus.DebugLevel)
}

func TestRegister(t *testing.T) {
	err := metro2.Register(config, metro2.Rand32(), "1361170304012312", "13611703040")
	if err != nil {
		t.Error(err)
		return
	}
}

func TestAssign(t *testing.T) {
	data, err := metro2.Assign(config, metro2.Rand32(), "13611703040101", "13611703040", "13611703040", "6d6f24f3c3604fed886d5b8ec3c8779d")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(data)
	t.Log("success.......")
}

func TestH5Detail(t *testing.T) {
	h5url := metro2.H5Detail(config, "token", "xx")
	t.Log(h5url)
}
