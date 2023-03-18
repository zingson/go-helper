package hdecimal

import (
	"testing"
)

func TestFenToYuan(t *testing.T) {
	for i := 0; i < 1000; i++ {
		t.Log(FenToYuan(int64(i)))
	}
}

func TestFenToYuanString(t *testing.T) {
	for i := 0; i < 1000; i++ {
		t.Log(FenToYuanString(int64(i)))
	}
}
