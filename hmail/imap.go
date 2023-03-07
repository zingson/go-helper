package hmail

import (
	"fmt"
	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
	"github.com/emersion/go-message/charset"
	"github.com/emersion/go-message/mail"
	"golang.org/x/text/encoding/simplifiedchinese"
	"io"
	"io/ioutil"
	"log"
	"time"
)

func init() {
	charset.RegisterEncoding("GB2312", simplifiedchinese.HZGB2312)
	charset.RegisterEncoding("GB18030", simplifiedchinese.GB18030)
	charset.RegisterEncoding("GBK", simplifiedchinese.GBK)
}

type Criteria struct {
	Search *imap.SearchCriteria // 邮件查询条件
}

func ImapClient(auth *Auth) (c *client.Client, err error) {
	c, err = client.DialTLS(auth.Addr, nil) // Connect to server
	if err != nil {
		return
	}
	fmt.Println("Imap Client Connect to server")

	if err = c.Login(auth.Username, auth.Password); err != nil {
		return
	}
	fmt.Println("Logged in")
	return
}

func messageToMail(msg *imap.Message, section imap.BodySectionName) (item *ImapMail, err error) {
	item = new(ImapMail)

	if msg == nil {
		log.Panic("ERR_IMAP:Server didn't returned message")
	}

	r := msg.GetBody(&section)
	if r == nil {
		log.Panic("ERR_IMAP:Server didn't returned message body")
	}

	// Create a new mail reader
	mr, err := mail.CreateReader(r)
	if err != nil {
		log.Panic(err)
	}

	// Print some info about the message
	header := mr.Header
	var date time.Time
	if date, err = header.Date(); err == nil {
		item.Date = date
	}
	var from []*mail.Address
	if from, err = header.AddressList("From"); err == nil {
		for _, address := range from {
			item.From = address.Address
		}
	}
	var to []*mail.Address
	if to, err = header.AddressList("To"); err == nil {
		for _, address := range to {
			item.To = append(item.To, address.Address)
		}
	}
	var address []*mail.Address
	if address, err = header.AddressList("Cc"); err == nil {
		for _, addr := range address {
			item.Cc = append(item.Cc, addr.Address)
		}
	}
	var subject string
	if subject, err = header.Subject(); err == nil {
		item.Subject = subject
	}

	// Process each message's part
	for {
		var part *mail.Part
		part, err = mr.NextPart()
		if err == io.EOF {
			err = nil
			break
		}
		if err != nil {
			log.Panic(err)
		}
		switch h := part.Header.(type) {
		case *mail.InlineHeader: // This is the message's text (can be plain-text or HTML)
			b, _ := ioutil.ReadAll(part.Body)
			item.Html = string(b)
		case *mail.AttachmentHeader: // This is an attachment
			filename, _ := h.Filename()
			ctype, _, _ := h.ContentType()
			fileBytes, _ := ioutil.ReadAll(part.Body)
			item.Attach = append(item.Attach, &Attach{
				FileName:    filename,
				ContentType: ctype,
				Bytes:       fileBytes,
			})
		}
	}
	return
}
