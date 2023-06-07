package sm

import (
	"encoding/base64"
	"fmt"
	"github.com/tjfoc/gmsm/sm3"
)

// Hash SM3
func Hash(b []byte) []byte {
	return sm3.Sm3Sum(b)
}

// HashHex SM3 输出16进制格式
func HashHex(b []byte) string {
	return fmt.Sprintf("%x", Hash(b))
}

// HashBase64 SM3 输出Base64格式
func HashBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(Hash(b))
}
