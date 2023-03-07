package v3

//CallbackSet 设置消息通知地址API
// 文档：https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter9_1_12.shtml
func (mkt *MktService) CallbackSet(notifyUrl string) (result *CallbackSetResult, err error) {
	params := map[string]interface{}{
		"mchid":      mkt.Mchid,
		"notify_url": notifyUrl,
	}
	err = mkt.HttpPost("/v3/marketing/favor/callbacks", params, &result)
	return
}

type CallbackSetResult struct {
	UpdateTime Time   `json:"update_time"`
	NotifyUrl  string `json:"notify_url"`
}
