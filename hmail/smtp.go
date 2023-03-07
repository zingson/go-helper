package hmail

import (
	"bytes"
	"crypto/tls"
	"github.com/jordan-wright/email"
	"github.com/sirupsen/logrus"
	"mime"
	"net/smtp"
	"strings"
)

// SmtpSend 发送邮件
func SmtpSend(m *Mail, a *Auth) (err error) {
	e := email.NewEmail()
	e.From = m.From + " <" + a.Username + ">"
	e.To = m.To
	e.Cc = m.Cc
	e.Bcc = m.Bcc
	e.Subject = m.Subject
	e.Text = []byte(m.Text)
	e.HTML = []byte(m.Html) // Text 与 Html 二选一，都存在只发送Html

	if m.Attach != nil && len(m.Attach) > 0 {
		for _, attach := range m.Attach {
			_, _ = e.Attach(bytes.NewReader(attach.Bytes), mime.BEncoding.Encode("UTF-8", attach.FileName), attach.ContentType)
		}
	}
	host := strings.Split(a.Addr, ":")[0]

	//err = e.Send(s.Addr, smtp.PlainAuth("", s.Username, s.Password, host))
	err = e.SendWithTLS(a.Addr, smtp.PlainAuth("", a.Username, a.Password, host), &tls.Config{ServerName: host})
	if err != nil {
		return
	}
	logrus.Infof("SMTP发送邮件 %s   发件人：%s   收件人:%s", m.Subject, m.From, strings.Join(append(m.To, m.Cc...), ";"))
	return
}
