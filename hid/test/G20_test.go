package test

import (
	"github.com/zingson/go-helper/hid"
	"testing"
	"time"
)

func TestG20(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Log(hid.G20())
	}
}

func TestG6(t *testing.T) {
	go func() {
		for i := 0; i < 10; i++ {
			t.Log(hid.G6())
		}
	}()
	go func() {
		for i := 0; i < 10; i++ {
			t.Log(hid.G6())
		}
	}()
	go func() {
		for i := 0; i < 10; i++ {
			t.Log(hid.G6())
		}
	}()
	time.Sleep(10 * time.Second)
}
