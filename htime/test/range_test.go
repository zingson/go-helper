package test

import (
	"github.com/zingson/go-helper/htime"
	"testing"
)

func TestTimeRangeDaysF(t *testing.T) {
	beg, end := htime.TimeRangeDaysF(1, htime.LayoutT14)
	t.Log(beg, " - ", end)

	beg, end = htime.TimeRangeWeekF(1, htime.LayoutT14)
	t.Log(beg, " - ", end)

	beg, end = htime.TimeRangeMonthF(1, htime.LayoutT14)
	t.Log(beg, " - ", end)

	beg, end = htime.TimeRangeYearF(1, htime.LayoutT14)
	t.Log(beg, " - ", end)
}
