package hmcode

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"strconv"
)

//Bit 模10隔位乘2加校验数算法，计算步骤如下:
//1. 从右边第1个数字开始每隔一位乘以2
//2. 把在步骤1中获得的乘积的各位数字与原号码中未乘2的各位数字相加。
//3. 把步骤2得到的总和从该值的下一个以零结尾的数中减去〔得数是总和个位数字的“10”的补数〕。如果在步骤2得到的总和是以0结尾的数（30，40等等），则校验数字是0。
//@params ori A区+B区原始码，返回校验位
//@return 返回校验位数字
func Bit(ori string) string {
	var sum int64
	for i, c := range ori {
		cInt := parseInt(c)
		if i%2 == 0 {
			sum += cInt
			continue
		}
		x := cInt * 2
		for _, y := range strconv.FormatInt(x, 10) {
			sum += parseInt(y)
		}
	}
	bit := 10 - sum%10
	if bit == 10 {
		bit = 0
	}
	return strconv.FormatInt(bit, 10)
}

func parseInt(v int32) int64 {
	i, e := strconv.ParseInt(string(v), 10, 64)
	if e != nil {
		err := errors.New("Error:券码不能存在0-9数字以外的字符, " + e.Error())
		log.Error(err)
		panic(err)
	}
	return i
}

//FillZero 字符串前补0 原始字符串，需要的长度（前面补0个数）
func FillZero(ori string, wantLen int) string {
	oriLen := len([]rune(ori))
	if oriLen < wantLen {
		zero := "000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"
		ori = zero[0:wantLen-oriLen] + ori
	}
	return ori
}
