package test

import (
	"github.com/zingson/go-helper/sdk/aliyunoss"
	"testing"
)

func TestParseURL(t *testing.T) {
	bucketPath, err := aliyunoss.ParseURL("http://oss.himkt.cn/images%2F20201225%2Fud.png?Expires=1608969666&OSSAccessKeyId=LTAI4GBRkBG5Sn5sZNDvbhAc&Signature=MWrQfxb%2FQRivgWqgKZ6QqvSGM1M%3D")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("bucketPath: ", bucketPath)
}

func TestSignURL(t *testing.T) {
	bucket, err := aliyunoss.NewBucket(getConfig())
	if err != nil {
		return
	}
	fileUrl, err := aliyunoss.SignURL(bucket, "http://oss.himkt.cn/images/20201225/ud.png")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("fileUrl: ", fileUrl)
}
