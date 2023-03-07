package metro2

type Config struct {
	Appid       string `json:"appid"`         // 应用标识，地铁分配
	ServiceUrl  string `json:"service_url"`   // 服务链接，地铁分配
	AesSecret   string `json:"aes_secret"`    // Aes密钥，地铁分配
	AesIv       string `json:"aes_iv"`        // Aes Iv，地铁分配
	MetroRsaPub string `json:"metro_rsa_pub"` // 地铁平台公钥，地铁分配
	H5url       string `json:"h5url"`         // 乘车界面页面链接，地铁分配
	RsaPri      string `json:"rsa_pri"`       // 私钥，自己保留
	RsaPub      string `json:"rsa_pub"`       // 公钥，提交给地铁
}
