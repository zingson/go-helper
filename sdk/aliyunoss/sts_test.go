package aliyunoss

import (
	"encoding/json"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/sts"
	"testing"
)

func TestSts1(t *testing.T) {

	//构建一个阿里云客户端, 用于发起请求。
	//构建阿里云客户端时，需要设置AccessKey ID和AccessKey Secret。
	client, err := sts.NewClientWithAccessKey("cn-shanghai", getConfig().AccessKeyID, getConfig().AccessKeySecret)

	//构建请求对象。
	request := sts.CreateAssumeRoleRequest()
	request.Scheme = "https"

	//设置参数。
	request.RoleArn = "acs:ram::1895485311061251:role/oss"
	request.RoleSessionName = "accesskeyoss"
	request.DurationSeconds = "3600"

	//发起请求，并得到响应。
	response, err := client.AssumeRole(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Printf("response is %#v\n", response)
}

func TestSts(t *testing.T) {
	skk, err := Sts(getConfig())
	if err != nil {
		t.Error(err)
		return
	}
	sbs, _ := json.Marshal(skk)
	fmt.Println(string(sbs))
}
