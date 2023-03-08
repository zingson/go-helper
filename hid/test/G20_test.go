package test

import (
	"github.com/zingson/go-helper/hid"
	"testing"
)

func TestG20(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Log(hid.G20())
	}
}
