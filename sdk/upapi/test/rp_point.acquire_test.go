package test

import (
	"root/src/sdk/upapi"
	"testing"
	"time"
)

//PointAcquire 赠送专享红包
func TestPointAcquire(t *testing.T) {
	err := upapi.PointAcquire(cfgtoml(), &upapi.PointAcquireParams{
		TransSeqId: upapi.Rand32(),
		TransTs:    time.Now().Format("20060102150405"),
		InsAcctId:  "P220427102525711",
		PointId:    "4122042628792805",
		PointAt:    "1",
		BusiInfo: &upapi.BusiInfo{
			CampaignId:   upapi.Rand32(),
			CampaignName: "接口调用测试",
		},
		TransDigest:  "发测试红包",
		AccEntityTp:  upapi.AETP_01,
		Mobile:       "13611703040",
		CardNo:       "",
		OpenId:       "",
		ValidBeginTs: "",
		ValidEndTs:   "",
		Remark:       "",
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("success........")
}
