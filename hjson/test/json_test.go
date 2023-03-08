package test

import (
	"github.com/zingson/go-helper/hjson"
	"testing"
)

func TestParse(t *testing.T) {
	data := "{\"a\":\"aaaa\"}"
	v, err := hjson.Parse[map[string]string](data)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(v)
	s := hjson.Stringify(v)
	t.Log(s)
}

type Ms struct {
	A string `json:"a"`
}

func (o *Ms) Json() string {
	return hjson.Stringify(o)
}

func TestConvert(t *testing.T) {
	ms, err := hjson.Convert[map[string]any, Ms](map[string]any{"a": "123"})
	if err != nil {
		return
	}

	t.Log(ms.Json())
}
