package aliyunoss

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func New(config Config) (*oss.Client, error) {
	return oss.New(config.Endpoint, config.AccessKeyID, config.AccessKeySecret, oss.UseCname(config.Cname))
}

func NewBucket(config Config) (*oss.Bucket, error) {
	client, err := New(config)
	if err != nil {
		return nil, err
	}
	return client.Bucket(config.Bucket)
}
