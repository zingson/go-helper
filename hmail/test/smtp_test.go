package test

import (
	"root/src/pkg/hmail"
	"testing"
)

/*
账号:zengs@himkt.cn
独立密码： f2FDZkmtjsuKGv6f     IMAP/SMTP测试收发邮件使用。
*/

func TestSmtpSend(t *testing.T) {
	err := hmail.SmtpSend(&hmail.Mail{
		From:    "SMTP",
		To:      []string{"zengs@himkt.cn"},
		Cc:      nil,
		Bcc:     nil,
		Subject: "测试邮件1-1-1",
		Text:    "text内容",
		Html:    "<p>html内容</p>",
		Attach:  nil,
	}, &hmail.Auth{
		Addr:     smtpAddr,
		Username: username,
		Password: password,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("发送成功...")
}
