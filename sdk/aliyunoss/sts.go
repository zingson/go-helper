package aliyunoss

import (
	"errors"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/sts"
)

//STS GO示例文档：https://help.aliyun.com/document_detail/184381.html?spm=a2c4g.11186623.6.897.132347c13MIY5O

// 前端配置项文档：https://help.aliyun.com/document_detail/64095.html?spm=a2c4g.11186623.6.1480.27504164QfIHrA

// 授权前端操作对象存储文件
type StsBucketAccessKey struct {
	Endpoint        string `json:"endpoint"` // OSS endpoint ，可以使用绑定的域名
	Secure          bool   `json:"secure"`   // true 使用https，false不使用
	Cname           bool   `json:"cname"`    // endpoint 是否是自定义域名 true：是 false：否
	Region          string `json:"region"`
	AccessKeyID     string `json:"accessKeyId"`     // AccessId
	AccessKeySecret string `json:"accessKeySecret"` // AccessKey
	StsToken        string `json:"stsToken"`
	Bucket          string `json:"bucket"`
}

// StsOSS 阿里云Sts授权
func Sts(config *Config) (skk *StsBucketAccessKey, err error) {
	if config.Sts == nil {
		err = errors.New("Aliyun Oss Sts 授权配置不能为空")
		return
	}

	//构建一个阿里云客户端, 用于发起请求。
	//构建阿里云客户端时，需要设置AccessKey ID和AccessKey Secret。
	client, err := sts.NewClientWithAccessKey(config.RegionId, config.AccessKeyID, config.AccessKeySecret)
	if err != nil {
		return
	}

	//构建请求对象。
	request := sts.CreateAssumeRoleRequest()
	request.Scheme = "https"

	//设置参数。
	request.RoleArn = config.Sts.RoleArn
	request.RoleSessionName = config.Sts.RoleSessionName
	request.DurationSeconds = "3600"

	//发起请求，并得到响应。
	resp, err := client.AssumeRole(request)
	if err != nil {
		return
	}
	skk = &StsBucketAccessKey{
		Endpoint:        config.Sts.Endpoint,
		Cname:           config.Sts.Cname,
		Region:          config.Sts.Region,
		AccessKeyID:     resp.Credentials.AccessKeyId,
		AccessKeySecret: resp.Credentials.AccessKeySecret,
		StsToken:        resp.Credentials.SecurityToken,
		Bucket:          config.Bucket,
	}
	return
}
