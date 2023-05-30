package hid

import (
	"math/rand"
	"strconv"
)

// G6 6位数编号
func G6() string {
	v := rand.Intn(1000000)
	if v < 100000 {
		v += 100000
	}
	return strconv.Itoa(v)
}
