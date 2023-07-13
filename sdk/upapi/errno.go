package upapi

import (
	"encoding/json"
	"fmt"
	"strings"
)

/*
返回码	描述
94	没有通过人像验证，请次日后再试
95	您没有通过人像验证
96	当前用户未开通人像照片认证
a10	不合法的backend_token，或已过期（参见6.1.1获取backendToken章节，重新获取backend_token）
a20	不合法的frontend_token，或已过期（参见6.1.2获取frontToken章节，重新获取front_token）
a31	不合法的授权code，或已过期（参见5.3系统对接步骤章节，参见常见问题解答）
BC2126	缓存中不包含此商户appId
BC2127	缓存中不包含此backendToken
BC2133	获取商户信息异常（检查appId参数信息，若是无感支付，并检查planId）
BC0024	非法请求（接口返回检查时间戳，upsdk报错检查时间戳和安全域名）
BC0025	验签失败（检查签名因子和签名方法）
N60005	cdhdUsrId为空
N62100	商户appId不能为空
N62101	backendToken不能为空
N62102	frontToken不能为空
N62103	scope不能为空
N62104	签名字段值不能为空
N62105	responseType应为固定值code
N62106	grantType应为固定值authoriz
N62107	随机字符串不能为空
N62108	时间戳不能为空
N62109	accessToken不能为空
N62110	openId不能为空
N62111	url不能为空
N62112	code不能为空
N62113	contractCode不能为空
N62114	planId不能为空
N62115	contractId不能为空
N62116	content不能为空
N62117	secret不能为空
N62118	actionType不能为空
N62119	channelNo不能为空
N62120	region不能为空
N62121	result不能为空
N62122	bizOrderId不能为空
N62123	bizType不能为空
N62124	merchantId不能为空
N62125	notifyUrl不能为空
N62126	缓存中不包含此商户appId
N62127	缓存中不包含此backendToken
N62128	接口名不能为空
N62129	relateId不能为空
N62130	trId不能为空
N62131	无权限访问此接口（请检查授权联登上送scope是否正确，咨询业务人员是否具有该接口调用权限）
N62132	不支持此访问域名（请检查申请对接云闪付开放平台的申请表中配置的域名与授权联登请求链接中的redirect_url是否一致）
N62134	退税号不能为空
N62135	用户授权时，没有提供手机号
N62136	用户没有手机号
N62137	获取签约商户信息异常（检查上送的appId和planId）
N62138	验证支付密码异常
N62142	待更新的缓存key不能为空
N62143	待更新的缓存expire不能为空
S52131	无权限访问此接口（请检查授权联登上送scope是否正确，咨询业务人员是否具有该接口调用权限）
S52136	用户没有手机号
S52172	接入方编码非法（请检查活动id/机构账户号是否和申请单填写一致）


GPUP40040000  原因1：积分红包账户没钱了；
*/

var (
	E00   = ErrNew("00", "ok")
	EA10  = ErrNew("a10", "不合法的backend_token，或已过期")
	EA20  = ErrNew("a20", "不合法的frontend_token，或已过期")
	EA31  = ErrNew("a31", "不合法的授权code，或已过期")
	E3023 = ErrNew("3023", "请您先注册云闪付APP，再返回本页面参与活动")

	// 积分接口错误码定义
	gpup = map[string]*Err{
		"E3023":     E3023,
		"GPUP07609": ErrNew("GPUP07609", "接入方余额不足"),
		"GPUP07008": ErrNew("GPUP07008", "积分调账业务异常"),
		"GPUP07003": ErrNew("GPUP07003", "积分查询业务异常"),
		"GPUP07400": ErrNew("GPUP07400", "积分查询为空"),
		"GPUP07610": ErrNew("GPUP07610", "接入方关联积分Id不匹配"),
		"GPUP07600": ErrNew("GPUP07600", "调用规则查询积分信息异常/积分ID不存在"),
		"GPUP07002": ErrNew("GPUP07002", "积分赠送业务异常"),
		"GPUP07001": ErrNew("GPUP07001", "参数异常"),
		"SG0001":    ErrNew("SG0001", "参数不合法,请确认【accEntityTp】是否正确"),

		// 领券错误码
		"GCUP00030": ErrNew("GCUP00030", "优惠券不存在"),                                           //{"resp":"GCUP00030","msg":"The coupon is not exist.","params":{}}
		"GCUP06038": ErrNew("GCUP06038", "超过限制次数"),                                           //Coupon download failed due to useId is limited.[GCUP06038]
		"GCUP06005": ErrNew("GCUP06005", "券未到可用时间"),                                          //The coupon is not start to download.[GCUP06005]
		"GCUP06006": ErrNew("GCUP06006", "券已过有效期"),                                           //The coupon is end to download.GCUP06006
		"GCUP06007": ErrNew("GCUP06007", "规则验证失败，请确认在有效时间内"),                                 //Coupon download rules match failed.[GCUP06007]
		"GCUP07056": ErrNew("GCUP07056", "Coupon download failed due to cityId is invalid."), //Coupon download failed due to cityId is invalid.[GCUP07056]
		"GCUP07053": ErrNew("GCUP07053", "未绑定活动指定银行卡，请先去云闪付绑卡"),                              //Coupon download failed due to no cardNo is invalid.[GCUP07053]
		"GCUP06045": ErrNew("GCUP06045", "该优惠券已达到领用上限"),                                      //Coupon download failed due to there has no coupon left.[GCUP06045]
		"GCUP07052": ErrNew("GCUP07052", "云闪付没有绑卡或所绑卡不同名"),                                   //Coupon download failed due to acct is not samename.[GCUP07052]
		"GCUP06036": ErrNew("GCUP06036", "卡号受限无法领券"),                                         //Coupon download failed due to cardNo is limited.[GCUP06036]
		"GCUP07058": ErrNew("GCUP07058", "黄名单用户不能参与，谢谢！"),                                    //Coupon download failed due to yellowNameList check failed.[GCUP07058]
		"GCUP07060": ErrNew("GCUP07060", "您不符合领券要求，请查看活动说明"),                                 //Coupon download failed due to userId check failed.[GCUP07060]
		"GCUP07050": ErrNew("GCUP07050", " 闪券”entityTp“参数配置不一致"),                             //The coupon entityTp is not match
		"GCUP07051": ErrNew("GCUP07051", "当前用户未实名认证"),
	}
)

type Err struct {
	Code string `json:"errno"`
	Msg  string `json:"error"`
}

func ErrNew(code, msg string) *Err {
	return &Err{Code: code, Msg: msg}
}

func NewF(err error) *Err {
	e := err.Error()
	i := strings.Index(e, ":")
	if i == -1 {
		return &Err{Code: "UPERR", Msg: e}
	}
	return &Err{Code: e[0:i], Msg: e[i+1:]}
}

func (e *Err) NewMsg(msg string) *Err {
	return ErrNew(e.Code, msg)
}

func (e *Err) NewMsgF(args ...interface{}) *Err {
	return ErrNew(e.Code, fmt.Sprintf(e.Msg, args...))
}

func (e *Err) Error() string {
	//return e.Msg + ".[" + e.Code + "]"
	return e.Code + ":" + e.Msg
}

func (e *Err) JsonMarshal() []byte {
	b, _ := json.Marshal(e)
	return b
}
