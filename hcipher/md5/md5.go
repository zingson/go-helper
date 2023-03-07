package md5

import (
	"crypto/md5"
	"fmt"
)

// Md5 32位16进制字符串
func Md5(v string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(v)))
}
