package example

import (
	"crypto"
	"crypto/md5"
	cryptorand "crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"sort"
	"strings"
)

func main() {

	//ysfThird.MemberPointAcquire()
	bm := make(map[string]string)
	bm["txnType"] = "01"
	bm["txnSubType"] = "01"
	bm["bizType"] = "000201"
	bm["channelType"] = "08"
	bm["accessType"] = "0"
	bm["orderId"] = "17003880709960300587100268223"
	bm["txnTime"] = "20231121154922"
	bm["txnAmt"] = "1"
	bm["currencyCode"] = "156"
	bm["reqReserved"] = "-"
	bm["backUrl"] = "https://msd.himkt.cn/gw/62vip/order/call/back"
	FrontTransReq(bm)
}

// 前台接口交易
func FrontTransReq(bm map[string]string) (url string, kv map[string]string, err error) {
	var (
		requestId = Rand32()
		reqBody   string
	)
	defer func() {
		log.Info(requestId, "ocwap前端请求地址", url)
		log.Info(requestId, "ocwap前端请求报文", reqBody)
	}()
	//url = cfg.BaseServiceUrl + "/gateway/api/frontTransReq.do"
	bm["version"] = "5.1.0"
	bm["encoding"] = "UTF-8"
	bm["signMethod"] = "01"
	bm["certId"] = "86842351990" //签名方式01需要上送
	bm["merId"] = "89833027372F284"

	// 签名字符串
	signString := RsaSignSortMap(bm)

	// 计算签名
	sign, err := RsaWithSha256Sign(signString)
	if err != nil {
		return
	}
	bm["signature"] = sign
	kv = bm

	requestBodyBytes, _ := json.Marshal(bm)
	reqBody = string(requestBodyBytes)
	fmt.Println(sign)
	return
}

var PrivateKey = "-----BEGIN 私钥-----\nMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQDmPzFJZmsvO1nLmQqPR+DroEgC7dvJGb7u16ZWcImsxddWm8RgZUibnoJbQpmMBueocttmWSaLy7r2woh5SShjaBilv18F0BXzBUUijk2/GWxP1HhoLBmATYIQmo9y+1xzLdFiRySTtV1iIWVsXLGuGpxi8P04VpZnIfimp+U2VpAzKLOQcBABUcRz6RK0ugIhDj3fV+d0z1lL17OGow5ZfcqjqLmOdoqMQerfq41659iTmx1nsHnIQpcGqKiWgv9LuBzqQqF/74QIBKVZdIyQS1rKf8SfXublOg2auHyUTmE+wdGRcJtE/dvldVggKC9vzjqdtk9qlCPbmsI8q4htAgMBAAECggEADToyJapZnAQH25d2MNcLR94UZP6vMrVk76m555tdfWx8O4CFXphS/IW1rXJ9CbtX0aiPjCiw+zUAD68LUhRu2o8fY15EIjxbFnQGDb8pHkJio49sP2fRKzrIrxWCCjQAkqleLwwxQL8Zl0qd3gTRPFRRCaoYkgT/bPtZopFgrKdehbs71jd3G1GuTJhAquKY/vGciV9qZTZ9PFsu+NjCbM7K4N08jyUMgpBBxu/yjiI1AJOg9NB5oQuy7bWUZ3fqLm0vSdxmL4kLRn7dlH0Cfb215SGJHk5L9r0Tu1vSOWv7iRIumVu+4SW/RZVpBOUiosYnGozI9ign6FWJTK/TjQKBgQDqXsU6Y4qfabdVEgLU/Iqck5CBxqrLuOoaoPAaplcguPA4KywAWSBE68pNwAX7mtuv906SUe9RRykINDtJF9L0DtcUIuM0cRPsWv5GVm606BulbvuFml+XJpbkXz9r4Mb9XcsE/Z4hgjEqGBgmLrodis2n1Y/Qf8YTcDcJch/PkwKBgQD7fwD6eHxdtE+8ITLr2DFPAnf9k4u9HFHeq2uNc3SSEVdm12ZkdgbPRkEPEctozQZwFnFKzMRS5ccVBAdfj0G2W3OmdJOTI7duo22U+bxGB4BOj29sPywrbuB11SvaBFxI0DBG7l4e2/umfhqWWbkq/aF1BTeYfLaraf/c0C5H/wKBgCpjW4505GRyjUrgiAyXZEZvVepy0P8pgYWxZ9uaDBYJ/MBb5ohlljWA48TuHubjmioYM2i5UuxDT4OBVtugaHHloTauHKOmc/ECgAyeo0Uq2Z0xFmKSiglfXFkPdPwA2lyGsypC6G+tCVFQrEVznAutyAY71Nfj1CI94GZf8qrrAoGBAOadz4fedCx466TXzJn35T9ORDmC0uDW1zwYdMeA16QZwsvLOVll4ViwU3g+GQhA+WZPLhQpMN3TVnsu9pPqCZfuxF4fCDQkzE/oDNRPLRYKSt2/Ka1ZQGeThMCIeW2yEe2y+pIQgr5JerlxaW/1VpbVfPiRkCmEMv18deNGAyqBAoGAMh8e4Pu8N0Is3dqvAuHFr/cTMGMsY6Bo9HZRPowvNDw41XL0n3t6T1gjI3Z1fusVmodWOQ3diOS8MnP4JyMiiOhbPVqPrtpmiwMWXnuaFo/foBLOW2bPMrCrAhsrAUWd0yRIxA3P3QEoe9LTY6T26xa3jLiZcBdtk4mSXLQkE7Y=\n-----END 私钥-----"

// 签名
// @params value 待签名字符串
// @params priKey 格式为pem格式私钥，通过银联pfx证书提取私钥，  对应的提取cer格式公钥上传到银联商户平台
func RsaWithSha256Sign(value string) (sign string, err error) {
	p, _ := pem.Decode([]byte(PrivateKey))
	//X509解码
	privateKey, err := x509.ParsePKCS8PrivateKey(p.Bytes)
	//if err != nil {
	//	return
	//}

	sha256hex := fmt.Sprintf("%x", sha256.Sum256([]byte(value)))
	hash := sha256.New()
	hash.Write([]byte(sha256hex))
	shaBytes := hash.Sum(nil)
	b, err := rsa.SignPKCS1v15(cryptorand.Reader, privateKey.(*rsa.PrivateKey), crypto.SHA256, shaBytes)
	if err != nil {
		return
	}
	sign = base64.StdEncoding.EncodeToString(b)
	return
}

// 拼接待签名字符串
func RsaSignSortMap(params map[string]string) string {
	var (
		buf     strings.Builder
		keyList []string
	)
	for k := range params {
		keyList = append(keyList, k)
	}
	sort.Strings(keyList)
	for _, k := range keyList {
		k = strings.TrimSpace(k)
		if "signature" == k || k == "" {
			continue
		}
		v := strings.TrimSpace(params[k])
		buf.WriteString(k)
		buf.WriteByte('=')
		buf.WriteString(v)
		buf.WriteByte('&')
	}
	s := buf.String()
	s = s[0 : len(s)-1]
	return s
}

// Rand32 使用crypto/rand 随机赋值byte数组， 然后md5返回32位十六进制字符串
func Rand32() string {
	b := make([]byte, 48)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return fmt.Sprintf("%x", md5.Sum(b))
}

type RespBody struct {
	Resp   string      `json:"resp"`
	Msg    string      `json:"msg"`
	Params interface{} `json:"params"`
}
