package test

import (
	"encoding/json"
	"github.com/zingson/go-helper/htype"
	"testing"
)

type User struct {
	Status htype.Bool `json:"status"`
}

func TestA(t *testing.T) {
	v := &User{Status: htype.True}
	b, _ := json.Marshal(v)
	t.Log(string(b))

	var obj User
	json.Unmarshal(b, &obj)

	t.Logf("%v", obj)
}
