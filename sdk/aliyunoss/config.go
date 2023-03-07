package aliyunoss

// 阿里云配置，从阿里云平台获取
type Config struct {
	Endpoint        string     `json:"endpoint"`          // OSS endpoint
	Cname           bool       `json:"cname"`             // 是否使用自定义域名
	Bucket          string     `json:"bucket"`            //
	AccessKeyID     string     `json:"access_key_id"`     // AccessId
	AccessKeySecret string     `json:"access_key_secret"` // AccessKey
	RegionId        string     `json:"region_id"`         // regionId
	Sts             *StsConfig `json:"sts"`
}

// STS 参数
type StsConfig struct {
	Region          string `json:"region"`            // 可用区，如 oss-cn-shanghai
	Endpoint        string `json:"endpoint"`          // sts endpoint 自定义域名
	Secure          bool   `json:"secure"`            // true 使用https，false不使用
	Cname           bool   `json:"cname"`             // endpoint 是否是自定义域名  true使用自定义域名，false是可用区拼接域名
	RoleArn         string `json:"role_arn"`          // STS需要 Arn，"acs:ram::1895485311061251:role/oss" ,使用sts获取StsToken时需要指定角色权限
	RoleSessionName string `json:"role_session_name"` // STS需要 自定义名字，可选
}
