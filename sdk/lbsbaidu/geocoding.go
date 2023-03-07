package lbsbaidu

import (
	"encoding/json"
	"fmt"
	"strings"
)

type RetCoordtype string

const (
	GCJ2LL RetCoordtype = "gcj02ll" //国测局坐标,高德坐标
	BD09LL RetCoordtype = "bd09ll"  //百度经纬度坐标，默认
	BD09MC RetCoordtype = "bd09mc"  //百度墨卡托坐标
)

// Geocoding 地理编码
// 文档：https://lbs.baidu.com/index.php?title=webapi/guide/webservice-geocoding
func Geocoding(c *Config, address string, opts ...*GeocodingOptions) (r *GeocodingResult, err error) {
	address = strings.ReplaceAll(address, "-", "")
	address = strings.ReplaceAll(address, " ", "")
	address = strings.ReplaceAll(address, "\n", "")
	address = strings.ReplaceAll(address, "\r", "")
	var url = fmt.Sprintf("https://api.map.baidu.com/geocoding/v3/?output=json&ak=%s&address=%s", c.Ak, address)
	for _, opt := range opts {
		if opt.City != "" {
			url = url + "&city=" + opt.City
		}
		if opt.RetCoordtype != "" {
			url = url + "&ret_coordtype=" + string(opt.RetCoordtype)
		}
	}

	rBytes, err := HttpGet(c, url)
	if err != nil {
		return
	}

	var res *GeocodingRes
	err = json.Unmarshal(rBytes, &res)
	if err != nil {
		return
	}
	if res.Status != 0 {
		err = lbsbdErr("Geocoding Status %d %s", res.Status, res.Message) // 错误码对应错误信息，请查看文档
		return
	}
	r = res.Result
	return
}

type GeocodingOptions struct {
	City         string
	RetCoordtype RetCoordtype
}

type GeocodingRes struct {
	Status  int              `json:"status"`
	Message string           `json:"message"`
	Result  *GeocodingResult `json:"result"`
}

type GeocodingResult struct {
	Location *Location `json:"location"`
	Precise  int       `json:"precise"` //位置的附加信息，是否精确查找。1为精确查找，即准确打点；0为不精确，即模糊打点。
}

type Location struct {
	Lat float64 `json:"lat"` //纬度值
	Lng float64 `json:"lng"` //经度值
}

func (o *GeocodingResult) Json() string {
	oBytes, _ := json.Marshal(o)
	return string(oBytes)
}
