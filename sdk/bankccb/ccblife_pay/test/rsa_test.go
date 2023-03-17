package test

import (
	_ "embed"
)

//go:embed .secret/config.json
var cstr string

/*
/upactccb?ccbParamSJ=WFhwajV4QlliUm43OXFOVlpneDF6NHNiYTZockVNckJlM3VIRmFSY3ZrNVpkT2tJMXRvWTFwdHpkbXYyS21CVlRNbCsyMzR4VUdHTgp2a0xVa3Rqanplc1ExUi9mZGc1N2txS3ZOaUxXZnJ4UmE0Smd5UHhiMVZ1M01kQkNpVTZpNkpTUnlJZHVWb25CR09VUnhyT0dCcW5jCk13SmF3ejhpZUNtVUkwSHhieUZ2QUc4M0JjeFRjOTdNQnJPZU9BR1NSTjJQUUtTTVkydmhSSEZMZEdrTUZqcnhHRkhoT3EyRGxJS0EKcTZ2QjBvc0dKQWlGK1pTaGpiOHNXT29YZ3MyckR1OTVRcUxNQitJYmJXRHpLakxCWUlpYkxKSVZIZDE4WnFIaTk3MmdnRmRUM2JDRwpNVFkzN0RpVEFiT003M203eWs4MkpqTHZzVzVvZlk5NzFGcHV3Z0pRcE42NW84dGhNelZYeTc4Nkp5K1U3cjgxeW9JRVNrWUt0TlBmCm5JL1ZLVzlHQTFCc01LS01pajF6Rjl4ZFFmKzRHeWRqMFUyak1LSEJnRDczRUNkQ0RlSXdwS0FiVHh0U3d3Zm1GenZrZXhzMmVObjUKZnBvU0hTQi9rN2krdC8xN1VvQzNyb1QwdzVWdFpPS3dPbUZDSkJlQ3hUQjAxejdiVDM3b2ltV0QK&cityid=330200&userCityId=310000&CITYID=330200&USERCITYID=310000
*/
/*
func getConfig() (conf *ccblife_auth.Config) {
	err := json.Unmarshal([]byte(cstr), &conf)
	if err != nil {
		return
	}
	return conf
}

func TestRsaDecode(t *testing.T) {
	conf := getConfig()

	//data := "WFhwajV4QlliUm43OXFOVlpneDF6NHNiYTZockVNckJlM3VIRmFSY3ZrNVpkT2tJMXRvWTFwdHpkbXYyS21CVlRNbCsyMzR4VUdHTgp2a0xVa3Rqanplc1ExUi9mZGc1N2txS3ZOaUxXZnJ4UmE0Smd5UHhiMVZ1M01kQkNpVTZpNkpTUnlJZHVWb25CR09VUnhyT0dCcW5jCk13SmF3ejhpZUNtVUkwSHhieUZ2QUc4M0JjeFRjOTdNQnJPZU9BR1NSTjJQUUtTTVkydmhSSEZMZEdrTUZqcnhHRkhoT3EyRGxJS0EKcTZ2QjBvc0dKQWlGK1pTaGpiOHNXT29YZ3MyckR1OTVRcUxNQitJYmJXRHpLakxCWUlpYkxKSVZIZDE4WnFIaTk3MmdnRmRUM2JDRwpNVFkzN0RpVEFiT003M203eWs4MkpqTHZzVzVvZlk5NzFGcHV3Z0pRcE42NW84dGhNelZYeTc4Nkp5K1U3cjgxeW9JRVNrWUt0TlBmCm5JL1ZLVzlHQTFCc01LS01pajF6Rjl4ZFFmKzRHeWRqMFUyak1LSEJnRDczRUNkQ0RlSXdwS0FiVHh0U3d3Zm1GenZrZXhzMmVObjUKZnBvU0hTQi9rN2krdC8xN1VvQzNyb1QwdzVWdFpPS3dPbUZDSkJlQ3hUQjAxejdiVDM3b2ltV0QK"
	data := "SC9vWHN1bUtSd3JDQkVNMlg3cUJoeUtWa2RqUmJPN2F4cUFBRmJ4eEQ3RXdSeVFYcFdjQmlpL0ExaitWb1p3dGZpY1ZKSXFaK3R6V0xLcGNTYU5pU1dNYXE4eTdWTE52UFdIRG9pTS90VFJVM3Z4a1l5YnpvMGZZbnBaK20yaWx1MVUwYit5U1I2RHpHZFlWYXl4akFRZlQwU3ovRmxvQjhab29oTHUrVVRxRjlCeVpGbVZDQytvOWZ4cjNJbklpZXhwSHZyMmJ1dFhYa0Z2OFhmY3NrcjhlckIwRnVGNDdLL2dhYkVMWm83RHN3c01ybjhHTjdSa3cyci8zZzVuMm1KdC9IM2NmZnVoUlVhR24xc0wyZ0NBeHl3bmtTQnN2RFVWNUlBYVhSWXJSUHQxTytTNjdRNUx4MzR3YWtFSjdqYnFKdU1hRUs4NEJPOEVqRTliZ2hueEdRcFNsd2x5T2tXTzBOWGZqQm9QMnY2Z0ZFV2dOMHFMQWxlbmpwajNmVWptSDBHUGFudkJ2eHljMDU4UzF3VG1wSG05cjRkclIrRDVYUDQ2ZlpRSWdIVUpHUG0xSy9COFVQQVBTeDRkL00ybUwvK1Y3djFhNlZmcFJHNnZRaXZxdG05cGUzSkdSb3hxY3FieERSSnAvMHY5OGJPUHdOb2tpRjF1eGI1TkM="

	param, err := ccblife_auth.ParseCcbParamSJ(conf, data)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(param.Json())

}*/
/*
func TestRsaEncode(t *testing.T) {
	conf := getConfig()

	data := "BGCOLOR=&userid=YSM202112210902914&mobile=13611703040&cityid=330200&userCityId=310000&orderid=&PLATFLOWNO=0000A2UNK1640064349000416&openid=&lgt=121.45578&ltt=31.275461&Usr_Name=&USERID=YSM202112210902914&MOBILE=13611703040&CITYID=330200&USERCITYID=310000&ORDERID=&OPENID=&LGT=121.45578&LTT=31.275461&timestamp=20220127114219&TIMESTAMP=1643254939767"

	v, err := RsaEncode(data, conf.PubKey)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(v)
}

func TestMchRsaPub1(t *testing.T) {
	conf := getConfig()
	mpk, err := hex.DecodeString(conf.MchPubKey)
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log(string(mpk))
}

func TestAscii(t *testing.T) {

	t.Log(strconv.QuoteToASCII("中文"))
}
*/
