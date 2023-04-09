package htime

import (
	"time"
)

// RangeDays 天时间范围  n 0=当天 -1=前一天 1=后一天
func RangeDays(n int) (beg, end time.Time) {
	y, m, d := time.Now().Date()
	return time.Date(y, m, d, 0, 0, 0, 0, time.Local).AddDate(0, 0, n),
		time.Date(y, m, d, 23, 59, 59, 0, time.Local).AddDate(0, 0, n)
}

func RangeDaysF(n int, f string) (beg, end string) {
	b, e := RangeDays(n)
	return b.Format(f), e.Format(f)
}

// TimeRangeWeek 周时间范围 n 0=本周 -1=前一周 1=后一周
func RangeWeek(n int) (beg, end time.Time) {
	ntime := time.Now()
	offset := int(time.Monday - ntime.Weekday())
	// 周日time.Sunday=0 做特殊判断
	if offset > 0 {
		offset = -6
	}
	year, month, day := ntime.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, time.Local).AddDate(0, 0, offset+7*n),
		time.Date(year, month, day, 23, 59, 59, 0, time.Local).AddDate(0, 0, offset+6+7*n)
}

func RangeWeekF(n int, f string) (beg, end string) {
	b, e := RangeWeek(n)
	return b.Format(f), e.Format(f)
}

// RangeMonth 月时间范围 n 0=本月 -1=前一月 1=后一月
func RangeMonth(n int) (beg, end time.Time) {
	y, m, _ := time.Now().Date()
	return time.Date(y, m, 1, 0, 0, 0, 0, time.Local).AddDate(0, n, 0),
		time.Date(y, m, 1, 23, 59, 59, 0, time.Local).AddDate(0, 1+n, -1)
}

func RangeMonthF(n int, f string) (beg, end string) {
	b, e := RangeMonth(n)
	return b.Format(f), e.Format(f)
}

// RangeYear 本年时间范围 n 0=本年 -1=前一年 1=后一年
func RangeYear(n int) (beg, end time.Time) {
	y, _, _ := time.Now().Date()
	return time.Date(y, 1, 1, 0, 0, 0, 0, time.Local).AddDate(n, 0, 0),
		time.Date(y, 12, 31, 23, 59, 59, 0, time.Local).AddDate(n, 0, 0)
}

func RangeYearF(n int, f string) (beg, end string) {
	b, e := RangeYear(n)
	return b.Format(f), e.Format(f)
}
