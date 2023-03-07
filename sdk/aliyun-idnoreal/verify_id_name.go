package aliyun_idnoreal

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// 阿里云市场，二要素实名认证
// 文档地址：https://market.aliyun.com/products/57124001/cmapi00035880.html
// curl -i -k -X POST 'https://dfidveri.market.alicloudapi.com/verify_id_name'  -H 'Authorization:APPCODE 你自己的AppCode' --data 'id_number=445122********33&name=%E9%BB%84%E5%A4%A7%E5%A4%A7'

// Certification 实名认证
func Certification(conf *Config, idno, name string) (err error) {
	var (
		nlog    = logrus.WithField("qid", Rand32())
		reqUrl  = conf.ServiceUrl + "/verify_id_name"
		reqBody string
		resBody string
		begTime = time.Now().UnixMilli()
	)
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
		var errMsg string
		if err != nil {
			errMsg = "异常信息：" + err.Error()
		}
		ms := fmt.Sprintf("%d", time.Now().UnixMilli()-begTime)
		nlog.WithField("ms", ms).Infof("hali_idnoreal 实名认证API 请求URL：%s  请求Body：%s  响应Body: %s  %s  耗时：%sms", reqUrl, reqBody, resBody, errMsg, ms)
	}()

	var body = make(url.Values)
	body.Add("id_number", idno)
	body.Add("name", name)
	body.Encode()
	reqBody = body.Encode()

	request, err := http.NewRequest(http.MethodPost, reqUrl, strings.NewReader(reqBody))
	if err != nil {
		return
	}
	request.Header.Set("Authorization", "APPCODE "+conf.AppCode)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return
	}
	resBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return
	}
	resBody = string(resBytes)

	var result *CertificationResult
	err = json.Unmarshal(resBytes, &result)
	if err != nil {
		return
	}
	if response.StatusCode != http.StatusOK {
		err = errors.New("实名验证不通过")
		return
	}
	if result.State != 1 {
		err = errors.New("实名信息不一致")
		return
	}
	return
}

type CertificationResult struct {
	RequestId     string `json:"request_id"`
	Status        string `json:"status"`
	State         int64  `json:"state"`          // 1, //返回值为 1 : 查询成功, 二要素一致。  返回值为 2 : 查询成功, 二要素不一致  返回值为 4：库无
	ResultMessage string `json:"result_message"` // 结果信息
}

func (c *CertificationResult) Json() string {
	b, _ := json.Marshal(c)
	return string(b)
}
