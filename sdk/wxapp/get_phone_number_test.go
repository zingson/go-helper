package wxapp

import (
	"testing"
)

func TestGetPhoneNumber(t *testing.T) {

	r, err := GetPhoneNumber("", "", "", "")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(r.JSON())

}
