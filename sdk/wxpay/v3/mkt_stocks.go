package v3

import (
	"fmt"
)

// Stocks 查询批次列表
// 条件查询批次列表API https://pay.weixin.qq.com/wiki/doc/apiv3/wxpay/marketing/convention/chapter3_4.shtml
// 通过此接口可查询多个批次的信息，包括批次的配置信息以及批次概况数据。
func (mkt *MktService) Stocks(offset, limit int, stock_creator_mchid string, create_start_time, create_end_time string, status string) (res *MarketingFavorStocksRes, err error) {
	err = mkt.HttpGet(fmt.Sprintf("/v3/marketing/favor/stocks?offset=%d&limit=%d&stock_creator_mchid=%s&create_start_time=%s&create_end_time=%s&status=%s", offset, limit, stock_creator_mchid, create_start_time, create_end_time, status), &res)
	return
}

type MarketingFavorStocksRes struct {
	Total_count int64                       `json:"total_count"`
	Offset      int64                       `json:"offset"`
	Limit       int64                       `json:"limit"`
	Data        []*MarketingFavorStocksData `json:"data"`
}

type MarketingFavorStocksData struct {
	Stock_id             string `json:"stock_id"`
	Stock_creator_mchid  string `json:"stock_creator_mchid"`
	Stock_name           string `json:"stock_name"`
	Status               string `json:"status"`
	Create_time          string `json:"create_time"`
	Description          string `json:"description"`
	Available_begin_time string `json:"available_begin_time"`
	Available_end_time   string `json:"available_end_time"`
	Distributed_coupons  int64  `json:"distributed_coupons"`
	No_cash              bool   `json:"no_cash"`
	Singleitem           bool   `json:"singleitem"`
	Stock_type           string `json:"stock_type"`
}
