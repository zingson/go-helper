package hid

import (
	"github.com/google/uuid"
	"strings"
)

// UUID32 不包含-字符串的uuid
func UUID32() string {
	return strings.ReplaceAll(uuid.New().String(), "-", "")
}
