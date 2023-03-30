package hconv

import "strconv"

/**
 * Int类型转换为其它类型
 */

// Int64ToString int64转换string
func Int64ToString(v int64) string {
	return strconv.FormatInt(v, 10)
}
