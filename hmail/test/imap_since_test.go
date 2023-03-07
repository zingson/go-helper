package test

import (
	"root/hmail"
	"testing"
	"time"
)

func TestImapFetchSince(t *testing.T) {

	// 查询一天内的邮件
	list, err := hmail.ImapFetchSince(&hmail.Auth{
		Addr:     imapAddr,
		Username: username,
		Password: password,
	}, time.Now().Local().AddDate(0, 0, -2), time.Now().Local().AddDate(0, 0, -1))
	if err != nil {
		t.Error(err)
		return
	}
	for _, mail := range list {
		t.Log("标题：", mail.Subject)
		if mail.Attach != nil && len(mail.Attach) > 0 {
			t.Log("附件：", mail.Attach[0].FileName)
		}
	}

}
