package test

import (
	"github.com/zingson/go-helper/sdk/upapi"
	"testing"
	"time"
)

// PointAcquire 赠送专享红包
func TestPointAcquire(t *testing.T) {
	err := upapi.PointAcquire(cfgtoml(), &upapi.PointAcquireParams{
		TransSeqId: upapi.Rand32(),
		TransTs:    time.Now().Format("20060102150405"),
		InsAcctId:  "P230129155337412",
		PointId:    "4121010417055723",
		PointAt:    "1",
		BusiInfo: &upapi.BusiInfo{
			CampaignId:   upapi.Rand32(),
			CampaignName: "阿拉红包",
		},
		TransDigest:  "阿拉红包",
		AccEntityTp:  upapi.AETP_01,
		Mobile:       "15067420750",
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
