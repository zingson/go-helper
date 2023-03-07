package sha

import (
	"crypto/sha256"
	"fmt"
)

// Sha256 输出64位16进制字符串
func Sha256(v string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(v)))
}
