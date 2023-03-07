package ldpush

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
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"math/rand"
	"net/http"
	"sort"
	"strings"
	"time"
)

// Post 请求接口
func Post(conf *Config, path string, params map[string]string) (rbytes []byte, err error) {
	var (
		nlog       = logrus.WithField("channel", "LD").WithField("path", path)
		serviceUrl = conf.ServiceUrl + path
		reqBody    string
		resBody    string
	)
	defer func() {
		nlog.Info("请求URL:", serviceUrl)
		nlog.Info("请求报文:", reqBody)
		nlog.Info("响应报文:", resBody)
		if err != nil {
			nlog.Error("响应异常:", err.Error())
		}
	}()

	if params == nil {
		err = errors.New("ERR_LD_PARAMS:params is nil")
		return
	}
	params["appId"] = conf.Appid
	params["nonce"] = Rand32()
	params["timestamp"] = time.Now().Format("20060102150405")

	// 签名
	sign, err := Sign(params, conf.RsaPriKey)
	if err != nil {
		return
	}
	params["sign"] = sign

	pbytes, err := json.Marshal(params)
	if err != nil {
		return
	}
	reqBody = string(pbytes)
	resp, err := http.Post(serviceUrl, "application/json", strings.NewReader(reqBody))
	if err != nil {
		return
	}
	if resp.StatusCode != http.StatusOK {
		err = errors.New("ERR_LD_HTTP:" + resp.Status)
		return
	}
	rbytes, err = io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	resBody = string(rbytes)
	return
}

// 计算签名
func Sign(params map[string]string, rsaPriKey string) (sign string, err error) {
	var (
		keys []string
		buf  strings.Builder
	)
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for i, k := range keys {
		if i != 0 {
			buf.WriteByte('&')
		}
		if params[k] == "" {
			continue
		}
		buf.WriteString(k)
		buf.WriteByte('=')
		buf.WriteString(params[k])
	}
	pstr := buf.String() // 待签名字符串

	return SHA256withRSA(pstr, rsaPriKey)
}

func SHA256withRSA(value, priKey string) (sign string, err error) {
	var der []byte
	if strings.Contains(priKey, "BEGIN PRIVATE KEY") {
		p, _ := pem.Decode([]byte(priKey))
		if p == nil {
			err = errors.New("priKey pem.Decode error ")
			return
		}
		der = p.Bytes
	} else {
		der, err = base64.StdEncoding.DecodeString(priKey)
		if err != nil {
			err = errors.New("priKey Base64 Parse error " + err.Error())
			return
		}
	}
	privateKey, err := x509.ParsePKCS8PrivateKey(der)
	if err != nil {
		return
	}

	hash := sha256.New()
	hash.Write([]byte(value))
	shaBytes := hash.Sum(nil)
	b, err := rsa.SignPKCS1v15(cryptorand.Reader, privateKey.(*rsa.PrivateKey), crypto.SHA256, shaBytes)
	if err != nil {
		return
	}
	sign = base64.StdEncoding.EncodeToString(b)
	return
}

// Rand32 使用crypto/rand 随机赋值byte数组， 然后md5返回32位十六进制字符串
func Rand32() string {
	var b = make([]byte, 48)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return fmt.Sprintf("%x", md5.Sum(b))
}
