package test

import (
	"github.com/zingson/go-helper/hmongo"
	"testing"
)

func TestDsnLoad(t *testing.T) {
	t.Log(hmongo.Dsn())
}
