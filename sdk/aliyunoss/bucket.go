package aliyunoss

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"net/url"
)

//SignURL bucketPath = host+path
// 返回值为可直接访问的链接，有效期1小时
func SignURL(bucket *oss.Bucket, bucketPath string) (fileUrl string, err error) {
	if bucketPath == "" {
		return
	}
	p, err := url.Parse(bucketPath)
	if err != nil {
		return
	}
	fileUrl, err = bucket.SignURL(p.Path[1:], oss.HTTPGet, 3600)
	if err != nil {
		return
	}
	return
}

//ParseURL 参数为可访问的链接
// 返回存储格式链接
func ParseURL(fileUrl string) (bucketPath string, err error) {
	if fileUrl == "" {
		return
	}
	fileUrl, err = url.PathUnescape(fileUrl)
	if err != nil {
		return
	}
	p, err := url.Parse(fileUrl)
	if err != nil {
		return
	}
	bucketPath = p.Scheme + "://" + p.Host + p.Path
	return
}
