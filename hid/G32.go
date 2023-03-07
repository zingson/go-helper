package hid

import "time"

// G32 32位长度字符串
func G32() string {
	return time.Now().Local().Format("20060102150405") + UUID32()[14:]
}
