package sha

import (
	"crypto/sha1"
	"fmt"
)

// Sha1 输出40位16进制字符串
func Sha1(v string) string {
	return fmt.Sprintf("%x", sha1.Sum([]byte(v)))
}
