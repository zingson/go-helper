package hmcode

import (
	"errors"
	"strings"
)

/*
12位券码混淆排列说明，第五位固定为Y位数字，由Y决定使用哪一个组合
序号 0  1  2  3  4  5  6  7  8  9  10 11
标记 A0 A1 B0 B1 B2 B3 B4 B5 B6 B7 B8 Y
*/
var en12 = map[string][]string{
	"0": strings.Split("A0 A1 B7 B2 Y B4 B1 B8 B0 B6 B3 B5", " "),
	"1": strings.Split("A0 A1 B8 B4 Y B7 B0 B6 B1 B2 B5 B3", " "),
	"2": strings.Split("A0 A1 B3 B7 Y B2 B8 B1 B5 B0 B4 B6", " "),
	"3": strings.Split("A0 A1 B5 B8 Y B6 B2 B0 B3 B7 B1 B4", " "),
	"4": strings.Split("A0 A1 B2 B8 Y B0 B6 B5 B1 B4 B7 B3", " "),
	"5": strings.Split("A0 A1 B8 B5 Y B7 B2 B1 B4 B0 B3 B6", " "),
	"6": strings.Split("A0 A1 B7 B0 Y B4 B3 B1 B5 B2 B6 B8", " "),
	"7": strings.Split("A0 A1 B4 B3 Y B5 B2 B6 B1 B8 B0 B7", " "),
	"8": strings.Split("A0 A1 B8 B4 Y B6 B2 B5 B1 B3 B7 B0", " "),
	"9": strings.Split("A0 A1 B7 B6 Y B3 B1 B4 B2 B0 B5 B8", " "),
}

// k=标记 v=索引位置
var index12 = map[string]int8{"A0": 0, "A1": 1, "B0": 2, "B1": 3, "B2": 4, "B3": 5, "B4": 6, "B5": 7, "B6": 8, "B7": 9, "B8": 10, "Y": 11}

//Encode12 12位券码混淆编码
func Encode12(val string) (code string, err error) {
	if len(val) != 12 {
		err = errors.New("ERROR:required length must be 12 val=" + val)
		return
	}
	for _, v := range en12[val[11:12]] {
		i := index12[v]
		code += val[i : i+1]
	}
	return code, nil
}

//Decode12 12位券码混淆解码
func Decode12(code string) (val string, err error) {
	if len(code) != 12 {
		err = errors.New("ERROR:required length must be 12 val=" + val)
		return
	}
	var chars [12]string
	for i, v := range en12[code[4:5]] {
		chars[index12[v]] = code[i : i+1]
	}
	for _, c := range chars {
		val += c
	}
	return
}

// Bit12 11位数字+校验位 后混淆
func Bit12(s11 string) (string, error) {
	return Encode12(s11 + Bit(s11))
}

//Code12 计算加校验位与混淆的12位券码 ,A区2位 B区9位  校验码1位
// A区为固定值 B区为递增数字
func Code12(a2, b9 string) (string, error) {
	return Bit12(a2 + b9)
}

//Code12Verify 校验券码是否合法
func Code12Verify(coupons string) error {
	c, err := Decode12(coupons)
	if err != nil {
		return err
	}
	i := len(c) - 1
	if Bit(c[0:i]) != c[i:] {
		return ERR_INVALID_COUPON_CODE
	}
	return nil
}
