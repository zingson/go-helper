package test

import (
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io"
	"mime"
	"testing"
)

func TestMime(t *testing.T) {
	s0, _ := simplifiedchinese.GBK.NewEncoder().String("中文")
	s := mime.BEncoding.Encode("gbk", s0)
	t.Log(s)
	s = "=?GBK?B?xOPU2r2ty9XKodDeuMTBy9PKz+TD3MLr?="
	s2, err := (&mime.WordDecoder{
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
	}).Decode(s)
	if err != nil {
		panic(err)
	}
	t.Log(s2)
}
