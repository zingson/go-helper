package hlotto

import "testing"

func TestPrizeIdx(t *testing.T) {
	p := []string{"3", "2", "1"}
	PrizeIdx([]int64{20, 5, 3}, []int64{}, 1, func(idx, i int64) {
		println("idx=", idx, " 奖项", p[i])
	})
}
