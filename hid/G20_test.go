package hid

import (
	"testing"
)

func TestID(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Log(G20())
	}
}
