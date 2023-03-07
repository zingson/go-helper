package test

import (
	"root/src/sdk/upapi"
	"testing"
)

func TestTRIPLE_DES_ECB_PKCS5_Encode(t *testing.T) {
	v, err := upapi.Encode3DES("0bd0da388645f73e45f23497491f5e340bd0da388645f73e", "15306662241")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(v)
}
