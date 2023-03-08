package test

import (
	"fmt"
	"github.com/zingson/go-helper/sdk/upapi"
	"net/http"
	"testing"
)

// 5.8.6  抽奖（红包/票券）<qual.reduce>
func TestQualReduce(t *testing.T) {
	/*
	  "mchId": "100002",
	  "transNumber":"423456789039",
	  "qualNum":"3a5d4792-48c3-4416-809d-ada5da535f84",
	  "qualType": "mobile",
	  "qualValue":"13912300661",
	  "activityNumber":"1320200615282465"*/
	_, err := upapi.QualReduce(cfgtoml(), "423456789039", "1320200615282465", "3a5d4792-48c3-4416-809d-ada5da535f84", "mobile", "13611703040")
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log("TestQualReduce success..................")
}

func Test_Notify(t *testing.T) {

	// 示例数据：/coupon/notify.do?discountNum=1&appId=9e211304be4a46fdb7dff03f7a01b2ef&transSeqId=cf5f05a01bd74144bde0432f8d27a567&discountAt=0&orderAt=0&operaTp=04&discountName=%E5%86%9C%E4%B8%9A%E9%93%B6%E8%A1%8C%E4%BF%A1%E7%94%A8%E5%8D%A135%E5%85%83%E5%88%B8&entityId=GZYp%2F%2F2B%2BtgTx9qudT9UkyXuBSLUdGq3M3tjvadYaNDseAAgFGr9glXGCCf91dF0&nonceStr=8qe28DCbkztJWTBs&timestamp=1607325521&entityTp=03&traceId=cf5f05a01bd74144bde0432f8d27a567&discountId=3102020120253762&transTs=20201207151841&signature=HBt9kwRfeXgPybvgdvFYapYtEIIShAh%2BzdpyjdJv8QrakQul4FBtiVcorAsaK84TqGtAGVxkp6MerD0%2Fln%2FDS9ZPcebJsdU8Xga638GtGY8cFE0JEu1URsh7c3j6nhDMJW9RCo9X9gNhz3izYXtqiebc%2BBM0abwPicXLF3S5apkTiwvc018nScTSOLZV1JUx16sL2yFBnQbEJxJaz2v9GRn3aOAj6BmzxUspNSq5hY3NpR9%2FKClbjF3yBxO4fbTebITFp7FXoJUCu%2FlVI3SpM6t8nXk8kyHYP7YhTO2wwix4xTiktGkbKy4wjcX1d3ptahRP%2BslPDcPie5v7kG1JAQ%3D%3D

	http.Handle("/coupon/notify.do", upapi.CouponNotifyHandler(func(appid string) (cfg *upapi.Config, err error) {
		cfg = &upapi.Config{
			ServiceUrl:    "https://open.95516.com/open/access/1.0",
			Appid:         "9e211304be4a46fdb7dff03f7a01b2ef",
			Secret:        "acf33e66a8a54f9d977d4143e1f17bfe",
			SymmetricKey:  "bad5200bfe4a91e5cb02f1f2ef1aec08bad5200bfe4a91e5",
			UpPublicKey:   "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA0QRJ81dxUdJNXoJwx81d\nvExIWP9zGhVVdYWKgOajcQI/5F1Qt67ipEL+pSh30P9roPBv6LWHb42z/htmPUrK\nXJ4f/WspXkbfBZsERe8XT8NZRnSdR3iZ9RqJKMzgjOetuoeFzTQ5QBalQKfQN9g5\n8FEY0wrGH8DbrRzRImsnOVl0vvdIrqvTji+vD6GzZ8egSz9HZ0e9fQKG4dI1nuH1\n45OfHY/fNe23oWINbXfFpVWiw+WgTTf8XzjVERD3qAT4i3cwB8RdhNlk3ysW0EJr\nt2/WOJiI2NNK3xzXohqPYdUDRA4aWbRPtIma5EtBcnLFm76mXwkTlk9PJm7CJA3c\n2QIDAQAB\n-----END PUBLIC KEY-----\n",
			MchPrivateKey: "-----BEGIN PRIVATE KEY-----\r\nMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQC7kdKuAnMgu7AV\r\nD4hfaT9i4TpDBVxN0xA2A6vGppaHB5F8N9vHRCBJLdAm04+HkdDmxG41Hbq7lABn\r\nrM92vrCbSA1Lo6asOawaif8dnu/i92SOnoK+r7v0vNR+hfIAXqO2xCJB9a1IvprN\r\nUw9V8m8ALr8eLHBJ+0sFsnrgvJG4gYl4q0/+pzERJK0SKqFceUKpcqunfsZPv4ko\r\nXK9Q7kSZxi+i3iwwmQ/7IBnsEVB78v4KGrkXDI/nrydauyGGXYXkTCbFiOw8CuxA\r\nSkx4k0kwpA8nuvFXzBG0V7EMox9oQtiOkfGJgDUJvmYYqg8rpgDb54iAvKZMmz+4\r\nx/4UupJjAgMBAAECggEAFVVqnvwMWCbAykRwAFoaKYbwd3r+mqNs7pfQS9HawRTt\r\nSTGZP7rR6UDase/SHVtKZVTmLAhrmrYkraYMGrdpot+5E2dTp7cPih0z9QyEwE3f\r\nFBGXUVTvjdCEYredZMle2YTJWLM2uFVligDud5oRYfXvKuFnDCMWz1kTfMg10sRH\r\nAw5lLHJk1cJSsB+s8swHzch+IsFg5oyA6VcpFKPiKvMwy8m1A923nH+mVVVcj4wM\r\nB0/qCtxlmIAUUg5MKp/RGgKPPxReTg6bqF2t6wNrHZevVOsFhDStizwf6dwLYCLV\r\nwK3I6Szwp+7uR284hVKZLu2uwSuREdxi1Xc6cGMuAQKBgQDrWpVo0LKaLNjnbcVL\r\nUuFCyuN0M8aWMfMrG1NtoYYPC5hFK5ZTAV3euf7yTGHazfswJ9A6U+h9+0XsIhfY\r\n8lN9HvO2ah05Uo2EPSGhaLa0ziFw4Nbi5usd3X6vnQq270Q/BoHi8fSoDou9oelQ\r\nneG3zIZab3cHYzEzKnq1rbUEgQKBgQDMBiJCTKx94lU8UdIIRDg6VEfIrya3+9F4\r\nBTQI4xiSdikdz5iZC2tt5gGHnaKeYVwsAioHJmFBWyWu0YQXb0qt1X+vbM/d9v1z\r\nT/mJ3KSpx928RdkwVKQsGQFsYjPgDVpweQkzybFEJFstOaNHXIw3+RhQG37UXuOr\r\n6sD05B2U4wKBgQDongyEn5mXpvHvs+BH9a/te2japn4GX3JPzd9kwTwmTLiAzXbz\r\nrashA8cHpxUk1WgLDZ7St7JYKm3O2VemtsRsK5aIWlNuH7j91goSZdQH2qDU13Ws\r\nqL4EM7MOUfKQIubaQE1KiQjevhnCIXDgnFvHdV/prLgB1jl/r+G/BeSfgQKBgFqH\r\nfjwc+Y0CGQAi7ids3eZD73ZFAdExk8jFxkkLO6QBek0YCIYgYxLotFUQxU+xs8xz\r\nSWLSzOTLJPVlUk9zupdX3MhiZ/n91oiMPBXIKeiMHv+jnrOrWw2WKuOEz6/jPPYb\r\nPtIT9OxflXWD1ceccTuE9BzXlnd1g2CNUgFYFygxAoGAbM2zG5dJSR4icH6LqBoH\r\nEzgxZuLlRIyWS2wnyxoIneRfKEoXAnaeZTho0jYBqFsKsButRIVP1DLlzS99NCi0\r\nN2MTKkYWZZbs2kFin7sB92Xy3QYoHeru4fZK3MdBRj85e17n9MBfVCnGTq6cbxbm\r\nMG7BFjV7aoR/h2bkkVv6mxw=\r\n-----END PRIVATE KEY-----",
		}
		return
	}, func(req *http.Request, requestId string, notifyResult *upapi.CouponNotify) error {
		fmt.Println("通知结果：", notifyResult.JsonString())
		return nil
	}))
	http.ListenAndServe(":9031", nil)
}
