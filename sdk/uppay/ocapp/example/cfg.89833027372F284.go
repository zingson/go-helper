package example

import (
	"fmt"
	"github.com/zingson/go-helper/sdk/uppay/ocapp"
	"io/ioutil"
)

// 测试配置
var cfg89833027372F284 = ocapp.NewConfig("https://gateway.95516.com", "89833027372F284", "", "86842351990")

func init() {
	// 读取商户私钥, 此文件由商户通过银联平台下载的pfx证书导出，对应的公钥通过银联商户平台上传到银联
	bytes, err := ioutil.ReadFile("./.secret/89833027372F284.key")
	if err != nil {
		fmt.Println("Error: ", err.Error())
		return
	}
	cfg89833027372F284.MerPrivateKey = string(bytes)
}
