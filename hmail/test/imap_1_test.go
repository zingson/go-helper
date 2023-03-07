package test

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/axgle/mahonia"
	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"log"
	"mime"
	"net/mail"
	"testing"
	"time"
)

func getClient() *client.Client {
	// Connect to server
	c, err := client.DialTLS(imapAddr, nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected")

	// Don't forget to logout
	//defer c.Logout()

	// Login
	if err := c.Login(username, password); err != nil {
		log.Fatal(err)
	}
	log.Println("Logged in")
	return c
}

func TestImap1(t *testing.T) {
	log.Println("Connecting to server...")

	c := getClient()
	defer c.Logout()

	// List mailboxes
	/*mailboxes := make(chan *imap.MailboxInfo, 10)
	done := make(chan error, 1)
	go func() {
		done <- c.List("", "*", mailboxes)
	}()

	log.Println("Mailboxes:")
	for m := range mailboxes {
		log.Println("* " + m.Name)
	}

	if err := <-done; err != nil {
		log.Fatal(err)
	}*/

	// Select INBOX
	mbox, err := c.Select("INBOX", false)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Flags for INBOX:", mbox.Flags)

	// Get the last 4 messages
	from := uint32(1)
	to := mbox.Messages
	if mbox.Messages > 3 {
		// We're using unsigned integers here, only substract if the result is > 0
		from = mbox.Messages - 3
	}
	seqset := new(imap.SeqSet)
	seqset.AddRange(from, to)
	messages := make(chan *imap.Message, 10)
	done := make(chan error, 1)
	go func() {
		done <- c.Fetch(seqset, []imap.FetchItem{imap.FetchEnvelope}, messages)
	}()

	log.Println("Last 4 messages:")
	for msg := range messages {
		log.Println("* " + SubjectDecode(msg.Envelope.Subject))
	}

	if err := <-done; err != nil {
		log.Fatal(err)
	}

	log.Println("Done!")
}

func TestExampleClient_Fetch(t *testing.T) {
	// Let's assume c is a client
	var c *client.Client

	c = getClient()
	defer c.Logout()

	// Select INBOX
	mbox, err := c.Select("INBOX", false)
	if err != nil {
		log.Fatal(err)
	}

	// Get the last message
	if mbox.Messages == 0 {
		log.Fatal("No message in mailbox")
	}
	seqset := new(imap.SeqSet)
	seqset.AddRange(mbox.Messages, mbox.Messages)

	// Get the whole message body
	section := &imap.BodySectionName{}
	items := []imap.FetchItem{section.FetchItem()}

	messages := make(chan *imap.Message, 1)
	done := make(chan error, 1)
	go func() {
		done <- c.Fetch(seqset, items, messages)
	}()

	log.Println("Last message:")
	msg := <-messages
	r := msg.GetBody(section)
	if r == nil {
		log.Fatal("Server didn't returned message body")
	}

	if err := <-done; err != nil {
		log.Fatal(err)
	}

	m, err := mail.ReadMessage(r)
	if err != nil {
		log.Fatal(err)
	}

	header := m.Header
	log.Println("Date:", header.Get("Date"))
	from, _ := simplifiedchinese.GBK.NewDecoder().String(header.Get("From"))
	log.Println("From:", from)
	log.Println("To:", header.Get("To"))
	log.Println("Subject: ", SubjectDecode(header.Get("Subject")))

	body, err := ioutil.ReadAll(m.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Body: ")
	log.Println(body)
	log.Println(string(body))

	ubody := BodyDecode(string(body))

	log.Println(ubody)
}

// 任意编码转特定编码
func ConvertToStr(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}

func TestMain2(t *testing.T) {
	var c *client.Client
	var err error

	c = getClient()
	defer c.Logout()

	mbox, err := c.Select("INBOX", false)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(mbox.Name)
	log.Println("Login success!")
	criteria := imap.NewSearchCriteria()
	//criteria.WithoutFlags=[]string{imap.SeenFlag}
	timeSince := "2020-01-07 00:00:00"
	tpl := "2006-01-02 15:04:05"
	loc, _ := time.LoadLocation("Local")
	t1, _ := time.ParseInLocation(tpl, timeSince, loc)
	log.Println(t)
	criteria.Since = t1
	ids, err := c.Search(criteria)
	if err != nil {
		log.Fatal("Search:", err)
	}
	if len(ids) > 0 {
		log.Println("IDs found:", ids)
		seqset := new(imap.SeqSet)
		seqset.AddNum(ids...)
		sect := &imap.BodySectionName{}
		messages := make(chan *imap.Message, 10)
		done := make(chan error, 1)
		go func() {
			done <- c.Fetch(seqset, []imap.FetchItem{sect.FetchItem()}, messages)

		}()

		log.Println("Unseen messages:")

		//r,_:=regexp.Compile(`[\w\W]*?考勤[\w\W]*?`)
		dec := new(mime.WordDecoder)
		dec.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
			switch charset {
			case "gb2312":
				content, err := ioutil.ReadAll(input)
				if err != nil {
					return nil, err
				}
				//ret:=bytes.NewReader(content)
				//ret:=transform.NewReader(bytes.NewReader(content), simplifiedchinese.HZGB2312.NewEncoder())

				utf8str := ConvertToStr(string(content), "gbk", "utf-8")
				t := bytes.NewReader([]byte(utf8str))
				//ret:=utf8.DecodeRune(t)
				//log.Println(ret)
				return t, nil
			case "gb18030":
				content, err := ioutil.ReadAll(input)
				if err != nil {
					return nil, err
				}
				//ret:=bytes.NewReader(content)
				//ret:=transform.NewReader(bytes.NewReader(content), simplifiedchinese.HZGB2312.NewEncoder())

				utf8str := ConvertToStr(string(content), "gbk", "utf-8")
				t := bytes.NewReader([]byte(utf8str))
				//ret:=utf8.DecodeRune(t)
				//log.Println(ret)
				return t, nil
			default:
				return nil, fmt.Errorf("unhandle charset:%s", charset)

			}
		}
		for msg := range messages {
			//subject :=msg.Envelope.Subject
			//
			//log.Println("* "+r.FindString(subject))
			//log.Println("* ",msg.Envelope.Date)
			r := msg.GetBody(sect)
			m, err := mail.ReadMessage(r)
			if err != nil {
				log.Println(err)
			}
			header := m.Header
			log.Println("Date:", header.Get("Date"))
			//log.Println("From:",header.Get("From"))
			//log.Println("To:",header.Get("To"))

			log.Println("Sunject:", header.Get("Subject"))
			ret, err := dec.Decode(header.Get("Subject"))
			if err != nil {
				ret, _ = dec.DecodeHeader(header.Get("Subject"))
			}
			log.Println("Subject:", ret)
			rbytes, _ := io.ReadAll(m.Body)
			rrb, _ := base64.StdEncoding.DecodeString(string(rbytes))
			bstring, err := dec.Decode(string(rrb))
			if err != nil {
				bstring, _ = dec.DecodeHeader(string(rrb))
			}
			log.Println("CN Body: ", bstring)
		}
		if err := <-done; err != nil {
			log.Fatal(err)
		}
	}
	log.Println("Done!")
}

var wordDecoder = &mime.WordDecoder{
	CharsetReader: func(charset string, input io.Reader) (io.Reader, error) {
		switch charset {
		case "gb2312":
			return transform.NewReader(input, simplifiedchinese.HZGB2312.NewDecoder()), nil
		case "gbk":
			return transform.NewReader(input, simplifiedchinese.GBK.NewDecoder()), nil
		case "gb18030":
			return transform.NewReader(input, simplifiedchinese.GB18030.NewDecoder()), nil
		default:
			return nil, fmt.Errorf("unhandle charset:%s", charset)
		}
	},
}

// SubjectDecode 邮件标题解码
func SubjectDecode(subject string) string {
	subject, err := wordDecoder.Decode(subject)
	if err != nil {
		panic(err)
	}
	return subject
}

// BodyDecode 邮件内容解码
func BodyDecode(body string) string {
	bodyBytes, err := base64.StdEncoding.DecodeString(body)
	if err != nil {
		panic(err)
	}
	body, err = simplifiedchinese.GBK.NewDecoder().String(string(bodyBytes))
	if err != nil {
		panic(err)
	}
	return body
}
