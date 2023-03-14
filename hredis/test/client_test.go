package test

import (
	"context"
	_ "embed"
	"github.com/zingson/go-helper/hredis"
	"testing"
)

//go:embed .dsn
var dsn string

var client = hredis.NowClient(dsn)

func TestNew(t *testing.T) {

	err := client.Set(context.Background(), "k2", "k2", -1).Err()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("set success")
}
