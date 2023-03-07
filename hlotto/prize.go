package hlotto

// PrizeIdx 根据奖品投放数量，计算奖品队列
// a=投放数量数组 , 小奖品在前
// b=剩余数量数组
// no=序号，从1开始
// prize=奖项数组索引
func PrizeIdx(a []int64, b []int64, idx int64, prize func(idx, i int64)) {
	if len(b) == 0 {
		b = append(b, a...)
	}

	var i int64
	var suma = sum(a)
	var sumb = sum(b)
	if sumb == 0 {
		println("PrizeIdx 抽奖结束.")
		return
	}

	for bi, v := range b {
		if v < 1 {
			continue
		}
		blot := float64(v) / float64(sumb)
		alot := float64(a[bi]) / float64(suma)
		if blot >= alot {
			i = int64(bi)
			prize(idx, i) // 序号，中奖索引
			break
		}
	}

	b[i] = b[i] - 1
	PrizeIdx(a, b, idx+1, prize)
}

func sum(arr []int64) (v int64) {
	for _, item := range arr {
		v = v + item
	}
	return
}
