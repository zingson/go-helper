package htime

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

const (
	LayoutT14 = "20060102150405"
)

type T14 time.Time

func (nt T14) String() string {
	return time.Time(nt).Local().Format(LayoutT14)
}

// MarshalJSON	JSON编码
func (nt T14) MarshalJSON() ([]byte, error) {
	t := time.Time(nt).Local()

	if y := t.Year(); y < 0 || y >= 10000 {
		return nil, errors.New("Time.MarshalJSON: year outside of range [0,9999]")
	}

	b := make([]byte, 0, len(LayoutT14)+2)
	b = append(b, '"')
	if !t.IsZero() {
		b = t.AppendFormat(b, LayoutT14)
	}
	b = append(b, '"')
	return b, nil
}

// UnmarshalJSON JSON解码
func (nt *T14) UnmarshalJSON(data []byte) error {
	v := strings.Trim(string(data), "\"")
	if data == nil || v == "" {
		*nt = T14(time.Time{})
		return nil
	}
	t, err := time.Parse(LayoutT14, v)
	if err != nil {
		return fmt.Errorf("解析时间字符串'%s'出错", v)
	}
	*nt = T14(t.Local())
	return err
}

// NowF14 取当前格式化时间
func NowF14() string {
	return T14(time.Now()).String()
}

// ParseT14 解析时间
func ParseT14(t14 string) time.Time {
	t14Time, err := time.Parse(LayoutT14, t14)
	if err != nil {
		return time.Time{}
	}
	return t14Time
}
