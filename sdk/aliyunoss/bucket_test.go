package aliyunoss

import (
	"testing"
)

func TestParseURL(t *testing.T) {
	bucketPath, err := ParseURL("http://oss.himkt.cn/images%2F20201225%2Fud.png?Expires=1608969666&OSSAccessKeyId=LTAI4GBRkBG5Sn5sZNDvbhAc&Signature=MWrQfxb%2FQRivgWqgKZ6QqvSGM1M%3D")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("bucketPath: ", bucketPath)
}

func TestSignURL(t *testing.T) {
	bucket, err := NewBucket(getConfig())
	if err != nil {
		return
	}
	fileUrl, err := SignURL(bucket, "http://oss.himkt.cn/images/20201225/ud.png")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("fileUrl: ", fileUrl)
}
