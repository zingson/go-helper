package test

import (
	"github.com/zingson/go-helper/sdk/wxapp"
	"testing"
)

func TestUploadShippingInfo(t *testing.T) {
	wxapp.UploadShippingInfo(accessToken, &wxapp.UploadShippingInfoParams{
		OrderKey:      wxapp.OrderKey{},
		LogisticsType: 0,
		DeliveryMode:  0,
		ShippingList:  nil,
		UploadTime:    "",
		Payer:         wxapp.Payer{},
	})
}
