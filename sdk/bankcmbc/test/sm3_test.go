package test

import (
	"github.com/zingson/go-helper/sdk/bankcmbc"
	"testing"
)

func TestSM3Hex(t *testing.T) {
	t.Log(bankcmbc.SM3Hex([]byte("15906566836")))
}
