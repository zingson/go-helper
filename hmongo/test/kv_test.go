package test

import (
	"encoding/json"
	"github.com/zingson/go-helper/hmongo"
	"os"
	"testing"
)

func dsn() string {
	b, err := os.ReadFile("../.secret/hmongo")
	if err != nil {
		panic(err)
	}
	return string(b)
}

var kvClient = hmongo.NewKv(dsn(), "dsys_config")

func TestGet(t *testing.T) {
	value := hmongo.Get[hmongo.Option](kvClient, "mongo.himkt")
	b, _ := json.Marshal(value)
	t.Log(string(b))
}
