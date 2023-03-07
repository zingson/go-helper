package hmail

import (
	"errors"
	"github.com/emersion/go-imap"
	"log"
)

// ImapFetchLatest 收取最新N个邮件
func ImapFetchLatest(auth *Auth, num uint32) (list []*ImapMail, err error) {

	c, err := ImapClient(auth)
	if err != nil {
		return
	}
	defer c.Logout()

	mbox, err := c.Select("INBOX", false) // Select INBOX
	if err != nil {
		return
	}

	// Get the last message
	if mbox.Messages == 0 {
		err = errors.New("ERR_IMAP:No message in mailbox")
		return
	}
	from := uint32(1)
	to := mbox.Messages
	if mbox.Messages > num {
		// We're using unsigned integers here, only substract if the result is > 0
		from = mbox.Messages - num
	}

	seqSet := new(imap.SeqSet)
	//seqSet.AddNum(mbox.Messages)
	seqSet.AddRange(from, to)

	// Get the whole message body
	var section imap.BodySectionName

	messages := make(chan *imap.Message, num)
	go func() {
		if err = c.Fetch(seqSet, []imap.FetchItem{section.FetchItem()}, messages); err != nil {
			log.Panic(err)
		}
	}()

	for msg := range messages {
		var item *ImapMail
		item, err = messageToMail(msg, section)
		if err != nil {
			return
		}
		list = append(list, item)
	}
	return
}
