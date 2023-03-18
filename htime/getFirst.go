package htime

// deprecated
func GetFirst() (dayStr string) {
	t := NowF8()
	content := t[:len(t)-2]
	dayStr = content + "01"
	return
}
