package sdk_lbsaliyun

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func Geocoding(address string, opts ...*GeocodingOptions) (r *GeocodingResult, err error) {
	address = strings.ReplaceAll(address, "-", "")
	address = strings.ReplaceAll(address, " ", "")
	address = strings.ReplaceAll(address, "\n", "")
	address = strings.ReplaceAll(address, "\r", "")
	appCode := "d22d67e0e3fc4ca0bd48b35defa3c215"
	var url = fmt.Sprintf("http://geo.market.alicloudapi.com/v3/geocode/geo?output=JSON")

	req, _ := http.NewRequest("GET", url, nil)

	q := req.URL.Query()
	q.Add("address", address)
	for _, opt := range opts {
		if opt.City != "" {
			q.Add("city", opt.City)
		}
		if opt.Callback != "" {
			q.Add("callback", opt.Callback)
		}
		if opt.Batch != "" {
			q.Add("batch", opt.Batch)
		}

	}
	req.Header.Add("Authorization", "APPCODE "+appCode)
	req.URL.RawQuery = q.Encode()
	var resp *http.Response
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		return
	}

	body, _ := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	var res *GeocodingRes
	if len(body) == 0 {
		return
	}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return
	}
	if res.Status != "1" {
		err = lbsalErr("Geocoding status  %s %s.", res.Status, res.Info) // 错误码对应错误信息，请查看文档
		return
	}

	if len(res.Geocodes) == 0 {
		return
	}
	r = res.Geocodes[0]

	//defer resp.Body.Close()

	return
}

var LBSAL_ERR = errors.New("LBSAL_ERR")

func lbsalErr(msg string, a ...interface{}) error {
	return errors.New(LBSAL_ERR.Error() + ":" + fmt.Sprintf(msg, a...))
}

type GeocodingOptions struct {
	City string
	//Output   string
	Callback string
	Batch    string
}

type GeocodingResult struct {
	Location string `json:"location"`
}

type GeocodingRes struct {
	Status   string             `json:"status"`
	Info     string             `json:"info"`
	Geocodes []*GeocodingResult `json:"geocodes"`
}
