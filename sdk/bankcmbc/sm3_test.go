package bankcmbc

import "testing"

func TestSM3Hex(t *testing.T) {
	t.Log(SM3Hex([]byte("15906566836")))
}
