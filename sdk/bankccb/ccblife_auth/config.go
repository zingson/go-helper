package ccblife_auth

type Config struct {
	PlatformId string `json:"platform_id"` // 服务方编号
	PriKey     string `json:"pri_key"`     // 服务方私钥 解密用户授权信息
}
