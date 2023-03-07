package lbsbaidu

import (
	_ "embed"
	"net/url"
	"testing"
)

//go:embed .ak
var ak string // 百度地图ak

var config = &Config{
	Appid: "2147163214",
	Ak:    ak,
}

func TestA3(t *testing.T) {
	r, err := Geocoding(config, url.QueryEscape("上海市静安区延长中路 581号"), &GeocodingOptions{
		RetCoordtype: GCJ2LL,
	})
	if err != nil {
		t.Log(err.Error())
		return
	}
	t.Log(r.Json())
}

func TestGeoconv(t *testing.T) {
	lng, lat, err := Geoconv(config, "121.44928118140284,31.26889721864015", CONV_GCJ02, CONV_WGS84)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(lng, ",", lat)
}

func TestWGS84toGCJ02(t *testing.T) {
	//104.195397,35.86166
	t.Log(WGS84toGCJ02(104.195397, 35.86166))
}
