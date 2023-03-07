package wxapp

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"strings"
)

// https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/signature.html
/*
加密数据解密算法
接口如果涉及敏感数据（如wx.getUserInfo当中的 openId 和 unionId），接口的明文内容将不包含这些敏感数据。开发者如需要获取敏感数据，需要对接口返回的加密数据(encryptedData) 进行对称解密。 解密算法如下：

对称解密使用的算法为 AES-128-CBC，数据采用PKCS#7填充。
对称解密的目标密文为 Base64_Decode(encryptedData)。
对称解密秘钥 aeskey = Base64_Decode(session_key), aeskey 是16字节。
对称解密算法初始向量 为Base64_Decode(iv)，其中iv由数据接口返回。
微信官方提供了多种编程语言的示例代码（（点击下载）。每种语言类型的接口名字均一致。调用方式可以参照示例。

另外，为了应用能校验数据的有效性，会在敏感数据加上数据水印( watermark )
*/

// AES CBC PKCS7
// 微信小程序返回密文解密
func DecryptData(appid, sessionKey, encryptedData, iv string) (oriData []byte, err error) {
	// Base64 解码
	sessionKeyBytes, err := base64.StdEncoding.DecodeString(sessionKey)
	if err != nil {
		err = errors.New("sessionKey base64 解码错误")
		return
	}
	ivBytes, err := base64.StdEncoding.DecodeString(iv)
	if err != nil {
		err = errors.New("iv base64 解码错误")
		return
	}
	encryptedDataBytes, err := base64.StdEncoding.DecodeString(encryptedData)
	if err != nil {
		err = errors.New("encryptedData base64 解码错误")
		return
	}

	block, err := aes.NewCipher(sessionKeyBytes)
	if err != nil {
		err = errors.New("Aes Cipher Error " + err.Error())
		return
	}
	blockMode := cipher.NewCBCDecrypter(block, ivBytes)
	oriData = make([]byte, len(encryptedDataBytes))
	blockMode.CryptBlocks(oriData, encryptedDataBytes)
	oriData = PKCS7UnPadding(oriData)

	if !strings.Contains(string(oriData), appid) {
		err = errors.New("Illegal encryptedData")
		return
	}
	return
}

func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
