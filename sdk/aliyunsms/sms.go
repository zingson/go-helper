package aliyunsms

import (
	"encoding/json"
	"errors"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/sirupsen/logrus"
)

type Config struct {
	RegionId        string `json:"region_id"`         // 可留空
	AccessKeyId     string `json:"access_key_id"`     // 阿里云获取
	AccessKeySecret string `json:"access_key_secret"` // 阿里云获取
}

func Client(config Config) (*dysmsapi.Client, error) {
	return dysmsapi.NewClientWithAccessKey(config.RegionId, config.AccessKeyId, config.AccessKeySecret)
}

// Send 发送短信
func Send(config Config, rid, phoneNumbers, signName, templateCode string, params map[string]string) (err error) {
	var resBody string
	defer func() {
		errMsg := ""
		if err != nil {
			errMsg = "  异常: " + err.Error()
		}
		logrus.Infof("rid:%s 阿里云发送短信 accessKeyId：%s  Tel：%s  signName：%s  templateCode：%s   params：%s  resBody:%s  %s", rid, config.AccessKeyId, phoneNumbers, signName, templateCode, params, resBody, errMsg)
	}()

	templateParam := "{}" // 模板参数
	if params != nil {
		pBytes, _ := json.Marshal(params)
		templateParam = string(pBytes)
	}

	client, err := Client(config)
	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"
	request.PhoneNumbers = phoneNumbers
	request.SignName = signName
	request.TemplateCode = templateCode
	request.TemplateParam = templateParam
	request.OutId = rid
	response, err := client.SendSms(request)
	if err != nil {
		return err
	}
	resb, _ := json.Marshal(response)
	resBody = string(resb)
	if response.Code != "OK" {
		err = errors.New("短信发送失败 " + response.Code + " " + response.Message)
		return
	}
	return
}
