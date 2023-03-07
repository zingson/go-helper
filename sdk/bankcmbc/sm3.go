package bankcmbc

import (
	"encoding/base64"
	"fmt"
	"github.com/tjfoc/gmsm/sm3"
)

// SM3 Hash
func SM3(b []byte) []byte {
	return sm3.Sm3Sum(b)
}

// SM3Hex SM3 输出16进制格式
func SM3Hex(b []byte) string {
	return fmt.Sprintf("%x", SM3(b))
}

// SM3B64 SM3 输出Base64格式
func SM3B64(b []byte) string {
	return base64.StdEncoding.EncodeToString(SM3(b))
}
