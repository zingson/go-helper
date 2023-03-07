package banknbcb

import (
	"encoding/json"
	"errors"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"root/hid"
	"sort"
	"strconv"
	"strings"
	"time"
)

// POST 调用接口
func POST(c *Config, p map[string]interface{}, out interface{}) (err error) {
	var (
		requestId = hid.G32()
		ms        = time.Now().UnixMilli()
		nlog      = logrus.WithField("requestId", requestId)
	)
	if p == nil {
		p = make(map[string]interface{})
	}
	p["mchnt_cd"] = c.MchntCd

	defer func() {
		if err != nil {
			if !strings.HasPrefix(err.Error(), "NBCB") {
				err = errors.New("NBCB" + err.Error())
			}
			nlog.Infof("hnbcb 异常信息: %s  耗时：%dms", err.Error(), time.Now().UnixMilli()-ms)
		}
	}()

	// 加签
	p["sign"], err = RsaSign(sortStr(p), c.MchPriKey)
	if err != nil {
		return
	}
	pbytes, _ := json.Marshal(p)
	body := string(pbytes)

	nlog.Infof("hnbcb 请求地址: %s  tran_code=%s  易收宝商户号： %s ", c.ServiceUrl, p["tran_code"], c.MchntCd)
	nlog.Infof("hnbcb 请求报文: %s", body)

	resp, err := http.Post(c.ServiceUrl, "application/json", strings.NewReader(body))
	if err != nil {
		return
	}
	if resp.StatusCode != 200 {
		err = errors.New("NBCB_HTTP_STATUS_" + resp.Status)
		return
	}
	bbytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	nlog.Infof("hnbcb 响应报文: %s  耗时：%dms", string(bbytes), time.Now().UnixMilli()-ms)

	var rmap = make(map[string]interface{})
	err = json.Unmarshal(bbytes, &rmap)
	if err != nil {
		return
	}
	if rmap["return_code"] != "00" {
		err = errors.New("NBCB" + rmap["return_code"].(string) + ":" + rmap["return_msg"].(string))
		return
	}
	rSign := rmap["sign"].(string)
	rmap["sign"] = ""

	// 验签
	err = RsaVery(sortStr(rmap), rSign, c.NbcbPubKey)
	if err != nil {
		nlog.Error(err.Error())
		err = errors.New("NBCB_SIGN:验签失败")
		return
	}

	err = json.Unmarshal(bbytes, out)
	if err != nil {
		return
	}
	return
}

// 排序
func sortStr(p map[string]interface{}) (v string) {
	var keys []string
	for k, _ := range p {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, key := range keys {
		var item = p[key]
		var value = ""
		if item == "" {
			continue
		}
		switch item.(type) {
		case string:
			value = item.(string)
		case int:
			value = strconv.Itoa(item.(int))
		case int64:
			value = strconv.FormatInt(item.(int64), 10)
		case bool:
			value = strconv.FormatBool(item.(bool))
		}
		v = v + "&" + key + "=" + value
	}
	v = v[1:]
	return
}
