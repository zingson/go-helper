package test

import (
	"fmt"
	"github.com/zingson/go-helper/hmcode"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestEncode12(t *testing.T) {
	val := "121234567891"
	t.Log("原始val：", val)
	code, err := hmcode.Encode12(val)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(val[11:12])
	t.Log("code = " + code)
	t.Log("=====恢复混淆=====")
	r, err := hmcode.Decode12(code)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("val = " + r)
}

// 执行100万次的时间  不带io操作，1一千万次7秒
func TestEncode12w100(t *testing.T) {
	begTime := time.Now().UnixNano()

	rand.Seed(time.Now().UnixNano())
	begval := 100000000
	for i := 0; i < 1000000; i++ {
		_, err := hmcode.Encode12("12" + strconv.Itoa(begval+i) + strconv.Itoa(rand.Intn(10)))
		if err != nil {
			t.Error(err)
			return
		}
	}
	endTime := time.Now().UnixNano()
	t.Log("结束：", (endTime-begTime)/1e6)
	// 10w 带log = 3秒
	// 100w 带log = 25秒
	// 100我 不带log = 730ms
	// 1000w 不带log  = 6秒
	//结论： 性能在log，计算无性能问题
}

func TestCode12(t *testing.T) {
	begTime := time.Now().UnixNano()
	sum := 1000000
	for i := 0; i < sum; i++ {
		c, err := hmcode.Code12("10", strconv.Itoa(100000000+i))
		if err != nil {
			t.Error(err)
			return
		}
		t.Log(c)
	}
	endTime := time.Now().UnixNano()
	t.Log(fmt.Sprintf("混淆:%d次，耗时:%d ms", sum, (endTime-begTime)/1e6))
	// 混淆:10000次，耗时:9 ms
	// 混淆:1000000次，耗时:1043 ms
}

func TestCode12Verify(t *testing.T) {
	err := hmcode.Code12Verify("100941990995")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("正确")
}

// 测试遍历多少次才能猜出正确券码
func TestCode12For(t *testing.T) {
	sum := 10
	INVALID_COUPON_CODE := 0
	for i := 0; i < sum; i++ {
		err := hmcode.Code12Verify("10094192165" + strconv.Itoa(i))
		if err == nil {
			INVALID_COUPON_CODE++
		}
	}
	t.Log(fmt.Sprintf("找到有效券码个数：%d", INVALID_COUPON_CODE))
	// 后2位9个
	// 后3位90个
	// 后4位900个
	// 后5位9000个
}
