package v2

import "strconv"

var (
	ERR_SUCCESS     = Err(999, "success")
	ERR_SIGN        = Err(1001, "响应数据验签失败")
	ERR_RS_DECRYPTY = Err(1002, "卡密解密出错")
)

type ErrCode struct {
	Code int64
	Msg  string
}

func (e *ErrCode) Error() string {
	return e.Msg + ".[CHYS" + strconv.FormatInt(e.Code, 10) + "]"
}

func Err(code int64, msg string) *ErrCode {
	return &ErrCode{
		Code: code,
		Msg:  msg,
	}
}
