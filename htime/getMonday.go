package htime

import (
	"time"
)

func GetMonday() (thisMonday string) {
	now := time.Now()
	thisMonday = GetMondayOfWeek(now, "20060102")
	return
}

func GetMondayOfWeek(t time.Time, fmtStr string) (dayStr string) {
	dayObj := GetZeroTime(t)
	if t.Weekday() == time.Monday {
		//修改hour、min、sec = 0后格式化
		dayStr = dayObj.Format(fmtStr)
	} else {
		offset := int(time.Monday - t.Weekday())
		if offset > 0 {
			offset = -6
		}
		dayStr = dayObj.AddDate(0, 0, offset).Format(fmtStr)
	}
	return

}

func GetZeroTime(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}
