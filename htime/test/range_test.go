package test

import (
	"github.com/zingson/go-helper/htime"
	"testing"
)

func TestTimeRangeDaysF(t *testing.T) {
	beg, end := htime.RangeDaysF(1, htime.LayoutT14)
	t.Log(beg, " - ", end)

	beg, end = htime.RangeWeekF(1, htime.LayoutT14)
	t.Log(beg, " - ", end)

	beg, end = htime.RangeMonthF(1, htime.LayoutT14)
	t.Log(beg, " - ", end)

	beg, end = htime.RangeYearF(1, htime.LayoutT14)
	t.Log(beg, " - ", end)
}
