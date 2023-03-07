package lbsbaidu

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"time"
)

var LBSBD_ERR = errors.New("LBSBD_ERR")

func lbsbdErr(msg string, a ...interface{}) error {
	return errors.New(LBSBD_ERR.Error() + ":" + fmt.Sprintf(msg, a...))
}

//HttpGet 百度地图接口GET请求
func HttpGet(c *Config, url string) (rBytes []byte, err error) {
	var (
		qid   = Rand32
		msbeg = time.Now().UnixMilli()
	)
	defer func() {
		var errmsg = ""
		if err != nil {
			errmsg = "接口异常：" + err.Error()
		}
		logrus.WithField("qid", qid()).Infof("lbsbaidu 应用编号 %s 请求URL %s 响应报文: %s  耗时 %dms %s", c.Appid, url, string(rBytes), time.Now().UnixMilli()-msbeg, errmsg)
	}()

	resp, err := http.Get(url)
	if err != nil {
		return
	}
	if resp.StatusCode != 200 {
		err = lbsbdErr("HTTPStatus%d", resp.StatusCode)
		return
	}
	rBytes, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	return
}
