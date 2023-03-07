package v3

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

// 定义Time 与 JSON互转的格式
// RFC3339
type Time time.Time

func (nt Time) String() string {
	return time.Time(nt).Local().Format(time.RFC3339)
}

// MarshalJSON implements the json.Marshaler interface.
// The time is a quoted string in RFC 3339 format, with sub-second precision added if present.
func (nt Time) MarshalJSON() ([]byte, error) {
	t := time.Time(nt).Local()

	if y := t.Year(); y < 0 || y >= 10000 {
		// RFC 3339 is clear that years are 4 digits exactly.
		// See golang.org/issue/4556#c15 for more discussion.
		return nil, errors.New("Time.MarshalJSON: year outside of range [0,9999]")
	}

	b := make([]byte, 0, len(time.RFC3339)+2)
	b = append(b, '"')
	if !t.IsZero() {
		b = t.AppendFormat(b, time.RFC3339)
	}
	b = append(b, '"')
	return b, nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
// The time is expected to be a quoted string in RFC 3339 format.
func (nt *Time) UnmarshalJSON(data []byte) error {
	v := strings.Trim(string(data), "\"")
	if data == nil || v == "" {
		*nt = Time(time.Time{})
		return nil
	}
	t, err := time.Parse(time.RFC3339, v)
	if err != nil {
		return fmt.Errorf("解析时间字符串'%s'出错，请使用RFC3339协议格式，时间格式示例：2006-01-02T15:04:05+08:00", v)
	}
	*nt = Time(t.Local())
	return err
}
