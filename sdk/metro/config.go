package metro

type Config struct {
	ServiceUrl string `json:"serviceUrl"`
	AppId      string `json:"appId"`
	Secret     string `json:"secret"`    // 签名秘钥
	SecretAes  string `json:"secretAes"` // 手机号加密秘钥
	// 计次票二维码H5页面，参数 {code} 二维码地址  {sign} 手机号加密字符串
	//示例：https://itapdev.ucitymetro.com/appentry?path=/ticket/qrcode-nbhy/{code}&sign={sign}&appId={appId}
	QrCode string `json:"qrcode"`
}
