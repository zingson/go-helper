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
		InsAcctId:  "P230129155471199",
		PointId:    "4120060413537864",
		PointAt:    "6163",
		BusiInfo: &upapi.BusiInfo{
			CampaignId:   upapi.Rand32(),
			CampaignName: "营业员红包",
		},
		TransDigest:  "营业员红包",
		AccEntityTp:  upapi.AETP_01,
		Mobile:       "-18757436314",
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
