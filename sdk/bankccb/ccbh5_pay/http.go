package ccbh5_pay

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"strings"
	"time"
)

// Post 接口请求
func Post(conf *Config, serviceUrl string, CLD_TX_CODE string, CLD_BODY interface{}) (resBody string, err error) {
	var (
		rid     = Rand32()
		milli   = time.Now().UnixMilli()
		reqBody string
		reqEn   string
		resEn   string
	)
	defer func() {
		errMsg := ""
		if err != nil {
			errMsg = "异常：" + err.Error()
		}
		millisecond := fmt.Sprintf("%d", time.Now().UnixMilli()-milli)
		nlog := logrus.WithField("mchid", conf.PlatformId).WithField("millisecond", millisecond)
		nlog.Infof("rid:%s 建行生活 明文 请求URL：%s  请求报文：%s  响应报文：%s  %s  | %sms", rid, serviceUrl, reqBody, resBody, errMsg, millisecond)
		nlog.Infof("rid:%s 建行生活 密文 请求URL：%s  请求报文：%s  响应报文：%s  %s  | %sms", rid, serviceUrl, reqEn, resEn, errMsg, millisecond)
	}()

	cntmap := map[string]interface{}{
		"CLD_HEADER": map[string]interface{}{
			"CLD_TX_CHNL": conf.PlatformId,
			"CLD_TX_TIME": time.Now().Local().Format("20060102150405"),
			"CLD_TX_CODE": CLD_TX_CODE,
			"CLD_TX_SEQ":  rid,
		},
		"CLD_BODY": CLD_BODY,
	}

	rbmapBytes, _ := json.Marshal(cntmap)
	reqBody = string(rbmapBytes)

	// 加密报文
	reqCnt, err := RsaEncode(reqBody, conf.PubKey)
	if err != nil {
		return
	}

	// 签名
	mac := strings.ToUpper(fmt.Sprintf("%x", md5.Sum([]byte(reqBody+conf.PriKey))))

	// 请求报文
	rmap := map[string]string{"cnt": reqCnt, "mac": mac, "svcid": conf.PlatformId}
	rmapBytes, _ := json.Marshal(rmap)
	reqEn = string(rmapBytes)

	// HTTP POST
	resp, err := http.Post(serviceUrl, "application/json", strings.NewReader(reqEn))
	if err != nil {
		return
	}
	rBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	resEn = string(rBytes)
	var resmap map[string]string
	err = json.Unmarshal(rBytes, &resmap)
	if err != nil {
		return
	}

	// 解密
	resBody, err = RsaDecode(resmap["cnt"], conf.PriKey)
	if err != nil {
		return
	}

	// 验签
	if strings.ToUpper(fmt.Sprintf("%x", md5.Sum([]byte(resBody+conf.PriKey)))) != resmap["mac"] {
		err = errors.New("签名错误")
		return
	}

	return
}

// Rand32 使用crypto/rand 随机赋值byte数组， 然后md5返回32位十六进制字符串
func Rand32() string {
	b := make([]byte, 48)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return fmt.Sprintf("%x", md5.Sum(b))
}
