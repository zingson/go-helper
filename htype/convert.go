package htype

import "strconv"

// Int64String int64转换string
func Int64String(data int64) string {
	var intString string
	intString = strconv.FormatInt(data, 10)
	return intString
}

// StringInt64 string转换int64
func StringInt64(data string) int64 {
	var stringInt64 int64
	if data != "" {
		stringInt64, _ = strconv.ParseInt(data, 10, 64)
	} else {
		stringInt64 = 0
	}
	return stringInt64
}

// Float64String float64转换string
func Float64String(data float64) string {
	var float64String string
	float64String = strconv.FormatFloat(data, 'f', -1, 64)
	return float64String
}

// StringFloat64 string转换float64
func StringFloat64(data string) float64 {
	var stringFloat64 float64
	if data != "" {
		stringFloat64, _ = strconv.ParseFloat(data, 64)
	} else {
		stringFloat64 = 0
	}
	return stringFloat64
}
