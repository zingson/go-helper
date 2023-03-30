package hconv

import "strconv"

/**
 * Int类型转换为其它类型
 */

// IntS int转string
func IntS(data int) string {
	return strconv.Itoa(data)
}

// Int64S int64转string
func Int64S(v int64) string {
	return strconv.FormatInt(v, 10)
}
