package hmail

import "time"

// Mail 发件信息
type Mail struct {
	From    string    // 发件名称，可为空
	To      []string  // 收件邮箱，必填
	Cc      []string  // 抄送，可选
	Bcc     []string  // 密送，可选
	Subject string    // 邮件标题，必填
	Text    string    // 邮件文本内容 Text 与 Html 二选一
	Html    string    // 邮件HTML内容
	Attach  []*Attach // 邮件附件
}

// ImapMail 收件信息
type ImapMail struct {
	Mail
	Date time.Time // 邮件时间
}

type Attach struct {
	FileName    string // 文件名
	ContentType string // 文件类型
	Bytes       []byte // 文件内容
}

type Auth struct {
	Addr     string `json:"addr"` // 使用SSL邮件服务器，发邮件是SMTP服务器，收邮件是IMAP服务器
	Username string `json:"username"`
	Password string `json:"password"`
}
