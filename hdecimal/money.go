package hdecimal

// FenToYuan 分转元2位小数
func FenToYuan(v int64) (val float64) {
	val, _ = NewFromInt(v).DivRound(NewFromFloat(100), 2).Float64()
	return
}

// FenToYuanString 分转元2位小数，返回字符串
func FenToYuanString(v int64) string {
	return NewFromInt(v).DivRound(NewFromFloat(100), 2).String()
}
