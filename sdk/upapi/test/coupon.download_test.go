package test

import (
	"encoding/json"
	"root/src/sdk/upapi"
	"testing"
	"time"
)

/*
测试活动ID：
云闪付删券活动ID: 1320200615282456   couponId:3102020072729846


活动ID1：1320200615282427
活动ID2：1320200615282448

        外部活动ID           之前给的立减活动ID（编码非法）     现在给的立减活动ID
满减券1  1320200615282456   3102020072429626            3102020072729846 （测试可以用）
满减券2  1320200615282465   3102020072429625            3102020072429635
*/
// 5.8.9  赠送优惠券 <coupon.download>
func TestCouponDownload(t *testing.T) {

	r, err := upapi.CouponDownload(cfgtoml(), &upapi.CouponDownloadParams{
		TransSeqId: upapi.Rand32(),
		TransTs:    time.Now().Format("20060102"),
		CouponId:   "-3102023011347995",
		CouponNum:  1,
		Mobile:     "17710046353",
	})
	if err != nil {
		t.Error(err.Error())
		return
	}
	b, _ := json.Marshal(r)
	t.Log(string(b))

}

/*13566569175
18248613990
18758306375
18658298350
15958373097*/

/*func TestCouponDownload2(t *testing.T) {
	tels := ""
	ts := strings.Split(tels, "\n")

	s := ""
	for _, tel := range ts {
		r, err := upapi.CouponDownload(cfgtoml(), &upapi.CouponDownloadParams{
			TransSeqId: upapi.Rand32(),
			TransTs:    time.Now().Format("20060102"),
			CouponId:   "3102022070782331",
			CouponNum:  1,
			Mobile:     tel,
		})
		if err != nil {
			t.Log("error msg->", err.Error())
			continue
		}
		s += tel + "," + r.TransSeqId + "\n"
		t.Log("success ", tel, " ", r.CouponId, " ", r.TransSeqId)
	}
	err := ioutil.WriteFile("fq3.txt", []byte(s), fs.ModeType)
	if err != nil {
		t.Log(err)
	}
}*/

/*
3102022072187196
3102022072187195
3102022072187194
3102022072187187
3102022072187196

*/
/*func TestCouponDownloadBatch(t *testing.T) {
	idstr := ""
	ids := strings.Split(idstr, "\n")
	for _, id := range ids {
		t.Log(id)
		r, err := upapi.CouponDownload(cfgtoml(), &upapi.CouponDownloadParams{
			TransSeqId: upapi.Rand32(),
			TransTs:    time.Now().Format("20060102"),
			CouponId:   id,
			CouponNum:  1,
			Mobile:     "13611703040",
		})
		if err != nil {
			t.Log("error ", id, " msg->", err.Error())
			continue
		}
		t.Log("success ", r.CouponId, " ", r.TransSeqId)
	}

}*/
