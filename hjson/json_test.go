package hjson

import (
	"testing"
)

func TestParse(t *testing.T) {
	data := "{\"a\":\"aaaa\"}"
	v, err := Parse[map[string]string](data)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(v)
	s := Stringify(v)
	t.Log(s)
}

type Ms struct {
	A string `json:"a"`
}

func (o *Ms) Json() string {
	return Stringify(o)
}

func TestConvert(t *testing.T) {
	ms, err := Convert[map[string]any, Ms](map[string]any{"a": "123"})
	if err != nil {
		return
	}

	t.Log(ms.Json())
}
