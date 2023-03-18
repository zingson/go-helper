package htime

// T14ToT19 时间转换
func T14ToT19(t14 string) (v string) {
	return ParseT14(t14).Format(LayoutT19)
}
