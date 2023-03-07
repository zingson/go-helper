package lbsbaidu

import (
	"encoding/json"
	"fmt"
)

type CoordsType int

const (
	CONV_WGS84  CoordsType = 1 //1：GPS标准坐标（wgs84）；
	CONV_GCJ02  CoordsType = 3 //3：火星坐标（gcj02），即高德地图、腾讯地图和MapABC等地图使用的坐标；
	CONV_BD09LL CoordsType = 5 //5：百度地图采用的经纬度坐标（bd09ll）；
	CONV_BD09MC CoordsType = 6 //6：百度地图采用的墨卡托平面坐标（bd09mc）;
)

// Geoconv 坐标转换
//文档： https://lbs.baidu.com/index.php?title=webapi/guide/changeposition
func Geoconv(c *Config, coords string, from, to CoordsType) (lng, lat float64, err error) {
	var url = fmt.Sprintf("https://api.map.baidu.com/geoconv/v1/?coords=%s&from=%d&to=%d&ak=%s", coords, from, to, c.Ak)
	rBytes, err := HttpGet(c, url)
	if err != nil {
		return
	}
	var r *GeoconvRes
	err = json.Unmarshal(rBytes, &r)
	if err != nil {
		return
	}
	if r.Status != 0 {
		err = lbsbdErr("Geoconv Status %d %s", r.Status, r.Message) // 错误码对应错误信息，请查看文档
		return
	}
	lng = r.Result.Lng
	lat = r.Result.Lat
	return
}

type GeoconvRes struct {
	Status  int            `json:"status"`
	Message string         `json:"message"`
	Result  *GeoconvResult `json:"result"`
}

type GeoconvResult struct {
	Lng float64 `json:"x"`
	Lat float64 `json:"y"`
}
