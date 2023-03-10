package test

import (
	"bytes"
	"crypto/tls"
	"github.com/jordan-wright/email"
	"github.com/zingson/go-helper/hmail"
	"mime"
	"net/smtp"
	"strings"
)

func NewSmtp() *Smtp {
	return &Smtp{Tls: true, To: []string{}, Cc: []string{}, Bcc: []string{}}
}

type Smtp struct {
	Tls      bool            // 是否使用TLS发邮件
	Addr     string          // SMTP服务，如 smtp.exmail.qq.com:465
	Username string          // 账号
	Password string          // 密码
	From     string          // 如：name <name@example.cn>
	To       []string        // 收件人
	Cc       []string        // 抄送
	Bcc      []string        // 密码抄送
	Subject  string          // 邮件标题
	Text     string          // 邮件文本内容
	Html     string          // 邮件HTML内容
	Attach   []*hmail.Attach // 邮件附件
}

func (s *Smtp) SetAddr(addr string) *Smtp {
	s.Addr = addr
	return s
}
func (s *Smtp) Auth(username, password string) *Smtp {
	s.Username = username
	s.Password = password
	return s
}

func (s *Smtp) SetForm(form string) *Smtp {
	s.From = form
	return s
}

func (s *Smtp) SetTo(to, cc, bcc []string) *Smtp {
	if to != nil {
		s.To = to
	}
	if cc != nil {
		s.Cc = cc
	}
	if bcc != nil {
		s.Bcc = bcc
	}
	return s
}

func (s *Smtp) SetSubject(subject string) *Smtp {
	s.Subject = subject
	return s
}

func (s *Smtp) SetText(text string) *Smtp {
	s.Text = text
	return s
}

func (s *Smtp) SetHtml(html string) *Smtp {
	s.Html = html
	return s
}

func (s *Smtp) SetAttach(attach []*hmail.Attach) *Smtp {
	s.Attach = attach
	return s
}

func (s *Smtp) PutAttach(attach *hmail.Attach) *Smtp {
	s.Attach = append(s.Attach, attach)
	return s
}

func (s *Smtp) Send() (err error) {
	host := strings.Split(s.Addr, ":")[0]

	e := email.NewEmail()
	e.From = s.From + " <" + s.Username + ">"
	e.To = s.To
	e.Cc = s.Cc
	e.Bcc = s.Bcc
	e.Subject = s.Subject
	e.Text = []byte(s.Text)
	e.HTML = []byte(s.Html)
	if s.Attach != nil && len(s.Attach) > 0 {
		for _, attach := range s.Attach {
			e.Attach(
				bytes.NewReader(attach.Bytes),
				mime.BEncoding.Encode("UTF-8", attach.FileName),
				attach.ContentType)
		}
	}
	if s.Tls {
		err = e.SendWithTLS(s.Addr, smtp.PlainAuth("", s.Username, s.Password, host), &tls.Config{ServerName: host})
	} else {
		err = e.Send(s.Addr, smtp.PlainAuth("", s.Username, s.Password, host))
	}
	if err != nil {
		return
	}
	return
}
