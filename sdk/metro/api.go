package metro

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
)

const (
	METHOD_POST      = "POST"
	METHOD_GET       = "GET"
	APPLICATION_JSON = "application/json"
	APPLICATION_FORM = "application/x-www-form-urlencoded"
)

// AuthByMobile 用户互认接口，根据手机号返回用户编号
func AuthByMobile(conf *Config, userPhone string) (userId string, err error) {
	bm := make(map[string]string)
	bm["userPhone"] = userPhone
	data, err := Request(conf, "/public/user/authbymobile", APPLICATION_FORM, METHOD_POST, bm)
	if err != nil {
		return
	}
	var dm = make(map[string]string)
	err = json.Unmarshal(data, &dm)
	if err != nil {
		return
	}
	userId = dm["userId"]
	return
}

// ProductInfo 获取计次票商品信息
func ProductInfo(conf *Config, productCode string) (prodInfo *ProdInfo, err error) {
	bm := make(map[string]string)
	bm["productCode"] = productCode
	data, err := Request(conf, "/public/monthlyticket/productinfo", APPLICATION_FORM, METHOD_GET, bm)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &prodInfo)
	if err != nil {
		return
	}
	return
}

type ProdInfo struct {
	AvailableTimes int64     `json:"availableTimes"` //可乘车次数
	ValidType      int64     `json:"validType"`      //有效期类型(0绝对，1相对)
	ValidStart     time.Time `json:"validStart"`     //绝对有效期-开始
	ValidEnd       time.Time `json:"validEnd"`       //绝对有效期-结束
	Duration       int64     `json:"duration"`       //相对有效期（天）
	SaleStatus     int64     `json:"saleStatus"`
	LimitBuyNumber int64     `json:"limitBuyNumber"`
}

// MonthlyTicketOpen 计次票开通接口
func MonthlyTicketOpen(conf *Config, userId, productCode, outOrderNo string) (data *TicketData, err error) {
	bm := make(map[string]string)
	bm["userId"] = userId
	bm["productCode"] = productCode
	bm["outOrderNo"] = outOrderNo
	dbytes, err := Request(conf, "/public/monthlyticket/open", APPLICATION_JSON, METHOD_POST, bm)
	if err != nil {
		return
	}
	err = json.Unmarshal(dbytes, &data)
	if err != nil {
		return
	}
	return
}

type TicketData struct {
	OrderNo     string               `json:"orderNo"` //票务扩展平台订单编号
	TicketCodes []string             `json:"ticketCodes"`
	Tickets     []*TicketDataTickets `json:"tickets"`
}
type TicketDataTickets struct {
	TicketCode      string `json:"ticketCode"`      //月票编号
	TicketTimes     int    `json:"ticketTimes"`     //月票次数
	TicketStartTime string `json:"ticketStartTime"` //月票有效起始时间，yyyy-MM-ddHH:mm:ss
	TicketEndTime   string `json:"ticketEndTime"`   //月票有效截止时间，yyyy-MM-ddHH:mm:ss
}

// 计次票核销推送

// Entry 计次票二维码H5页面嵌入
func Entry(conf *Config, code, mobile string) (h5 string, err error) {
	sign := base64.StdEncoding.EncodeToString(AesEncryptECB([]byte(mobile), []byte(conf.SecretAes)))
	h5 = conf.QrCode
	h5 = strings.ReplaceAll(h5, "{code}", code)
	h5 = strings.ReplaceAll(h5, "{sign}", url.QueryEscape(sign))
	h5 = strings.ReplaceAll(h5, "{appId}", conf.AppId)
	logrus.WithField("appid", conf.AppId).Info("URL ", h5, "  conf.SecretAes=", conf.SecretAes, " mobile=", mobile, " sign=", sign)
	return
}

// Sign 签名
func Sign(conf *Config, bm map[string]string) string {
	var (
		buf     strings.Builder
		keyList []string
	)
	for k := range bm {
		keyList = append(keyList, k)
	}
	sort.Strings(keyList)
	for _, k := range keyList {
		if "sign" == k {
			continue
		}
		if bm[k] == "" {
			continue // 不包含value为空的字段
		}
		buf.WriteString(k)
		buf.WriteByte('=')
		buf.WriteString(bm[k])
		buf.WriteByte('&')
	}
	buf.WriteString("key") // 拼接密钥的k为key，非文档写的直接拼接
	buf.WriteByte('=')
	buf.WriteString(conf.Secret)
	buf.WriteByte('&')

	s := buf.String()
	s = s[0 : len(s)-1]

	// md5
	sign := fmt.Sprintf("%x", md5.Sum([]byte(s)))
	return strings.ToUpper(sign)
}

// Request 表单Post请求
func Request(conf *Config, path string, contentType, method string, bm map[string]string) (data []byte, err error) {
	var (
		surl     = conf.ServiceUrl + path
		reqBody  string
		resBody  string
		begmilli = time.Now().UnixMilli()
	)
	defer func() {
		errMsg := ""
		if err != nil {
			errMsg = "\n异常: " + err.Error()
		}
		ms := fmt.Sprintf("%d", time.Now().UnixMilli()-begmilli)
		logrus.WithField("appid", conf.AppId).Infof("hmetro \n接口：%s  \n请求：%s  \n响应：%s  %s  \n耗时：%sms ", surl, reqBody, resBody, errMsg, ms)
	}()

	bm["appId"] = conf.AppId
	bm["nonceStr"] = Rand32()
	bm["sign"] = Sign(conf, bm)

	reqBytes, err := json.Marshal(bm)
	if err != nil {
		return
	}
	reqBody = string(reqBytes)

	fv := make(url.Values)
	for k, v := range bm {
		fv.Set(k, v)
	}

	var resp *http.Response
	if method == "POST" {
		if contentType == APPLICATION_JSON {
			resp, err = http.Post(surl, contentType, strings.NewReader(reqBody))
		} else {
			resp, err = http.PostForm(surl, fv)
		}
	}
	if method == "GET" {
		resp, err = http.Get(surl + "?" + fv.Encode())
	}
	if err != nil {
		return
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	resBody = string(bodyBytes)

	// 解析响应报文
	var rmap = make(map[string]interface{})
	err = json.Unmarshal(bodyBytes, &rmap)
	if err != nil {
		return
	}
	code := rmap["code"].(string)
	if code != "0000" {
		msg := ""
		if rmap["msg"] != nil {
			msg = rmap["msg"].(string)
		}
		err = errors.New(msg + ".[" + code + "]")
		return
	}
	data, err = json.Marshal(rmap["data"])
	if err != nil {
		return
	}
	return
}
