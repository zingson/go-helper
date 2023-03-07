package upapi

// 文档 https://opentools.95516.com/applet/#/docs/develop/oauth?id=_020202

/*
基础能力包	--	用户 openId
获取用户手机号信息	scope.mobile	手机号授权域
获取用户实名信息	scope.auth	实名授权域
银行卡能力	scope.bank	银行卡信息授权（仅限金融小程序使用）
人脸识别能力	scope.face	云闪付实名用户信息认证
*/

type Scope string

const (
	// 网页版
	UPAPI_BASE   Scope = "upapi_base"
	UPAPI_MOBILE Scope = "upapi_mobile"
	UPAPI_AUTH   Scope = "upapi_auth"

	// 小程序版
	SCOPE_BASE   Scope = "scope.base"
	SCOPE_MOBILE Scope = "scope.mobile"
	SCOPE_FACE   Scope = "scope.face"
)

//AccEntityTp 账户主体类型， 2 位，可选： 01 -手机号 02 -卡号 03 -用户（三选一）
type AccEntityTp string

const (
	AETP_01 AccEntityTp = "01" // 手机号
	AETP_02 AccEntityTp = "02" // 卡号
	AETP_03 AccEntityTp = "03" // 用户（卡号、手机号、openid 三选一）
)
