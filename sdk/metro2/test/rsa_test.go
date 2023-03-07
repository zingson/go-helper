package test

import (
	"root/src/sdk/metro2"
	"testing"
)

func TestVerify(t *testing.T) {
	pub := `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAwQ1hK/gT39XhFok2tWPs
6yugP3p9wOkughpcrg3OpwhvklBxQbk7Bq6YASMa/jWqgeb/xCDx1jEMtXbMpECl
WjaVsqLQ4VLVO+Lj0m4MTXJBdGz2yTtjr00CSLtMEgpSYRbsv2iWTCvROqErr3ba
kGt/zG3kmp/6cg5N/pm2CvBtrFK8kUPBmZHjz9uSC5LQ6eL9Hrb7E2qXSA0DEFcm
CyptYMqWxun5mNXo7+Ijc0dDCL+ewo4wcHSVP8bSmlOamhQiiFUPAJXWRpaPayiC
iJ0G0oHyNbp2ctwQyjGQkpslO+T3PKjeXgfv/daQg88yxX4M3dRNjUf0M9UHSEqo
BQIDAQAB
-----END PUBLIC KEY-----`
	sign := `vugXoS6Pqg80pkkUl3PBzvIgY3I88Z86pNHxvt2YuPgwOP6zDeogILFBWPf6bhire3HzENexy/dSNu5kcNpGvzdpKoFhlQjUJLN6PU/lAinepwC0jhOhIzTvL5ocQT2Iqrm5bcFEPs/iSIeIhpX8BjdMsQCljIo9kFmNjz4TvM9j14xXCJk6wV3klh2ruLDRLbgbvUJysxp0nJCS9Pge/Y+7uwV3WTm4p5o1IBkuy6UyrlNM+dXUSRLN2ZfRnxERGviBAn7wt9XHQH6rMIljasGTHSZlFa5UYRgRq41S1jVEYR8iZ7dUA/1XKMZHNeTR1vlvuf8Oley+OZK3iAo6Vw==`
	value := `appid=UC2202000001&body={"GuidUser":"abcdef1234567890"}&nonce=123456&timestamp=123456&version=2.0`
	err := metro2.RsaVerify(sign, value, pub)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("success.......")

}

func TestAes(t *testing.T) {
	key := "868971231616403394817a2360c4e8b2"
	iv := "8689712316164033"
	ciphertext, err := metro2.EncodeCBC("AES测试文本", key, "8689712316164033")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(ciphertext)
	// 预期：+0CQ8nuTf8eprUc3e2P7gQ==

	ciphertext = "+0CQ8nuTf8eprUc3e2P7gQ=="
	plaintext, err := metro2.DecodeCBC(ciphertext, key, iv)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(plaintext)
}
