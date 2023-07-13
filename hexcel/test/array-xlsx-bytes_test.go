package test

import (
	"github.com/zingson/go-helper/hexcel"
	"os"
	"testing"
)

func TestArrayXlsxBytes(t *testing.T) {

	b, err := hexcel.ArrayToXlsx([][]any{{"列1", "列2"}})
	if err != nil {
		t.Error(err)
		return
	}
	os.WriteFile("1.xlsx", b, os.ModeType)
}
