package sha

import (
	"crypto/sha512"
	"fmt"
)

// Sha512 输出128位16进制字符串
func Sha512(v string) string {
	return fmt.Sprintf("%x", sha512.Sum512([]byte(v)))
}
