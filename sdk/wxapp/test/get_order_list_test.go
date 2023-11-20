package test

import (
	"github.com/zingson/go-helper/sdk/wxapp"
	"testing"
)

func TestGetOrderList(t *testing.T) {

	wxapp.GetOrderList(accessToken, &wxapp.GetOrderListParams{
		PayTimeRange: nil,
		OrderState:   1,
		Openid:       "",
		PageSize:     10,
	})
}
