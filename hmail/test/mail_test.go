package test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// 邮箱测试账号
var (
	imapAddr string
	smtpAddr string
	username string
	password string
)

func init() {
	dbytes, err := ioutil.ReadFile("./mail_test.json")
	if err != nil {
		panic(err)
	}
	var acc map[string]string
	err = json.Unmarshal(dbytes, &acc)
	if err != nil {
		panic(err)
	}
	fmt.Println(acc)
	imapAddr = acc["imap"] //接收服务器  imap.exmail.qq.com:993  (使用SSL)
	smtpAddr = acc["smtp"] //发送服务器  smtp.exmail.qq.com:465  (使用SSL)
	username = acc["username"]
	password = acc["password"]
}
