package v2

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"strings"
)

/**
报文加密（AES）+密钥加密（RSA）方式方法如下：

加密步骤：
1. 【生成秘钥】生成随机32位字符串s1作为AES秘钥；
2. 【加密密钥】使用RSA方式加密AES密钥输出base64格式字符串s2
3. 【加密报文】AES加密报文输出base64格式字符串s3
4. 【加密结果】使用冒号拼接结果 ciphertext = s2 + ":" + s3 ；

解密步骤：
1. 【拆解密文】根据冒号拆分密文并做baes64解码，密文报文s3，密文AES密钥s2；
2. 【解密密钥】使用RSA解密密钥s2获得AES密钥s1；
3. 【解密报文】通过密钥s1使用AES解密密文s3；
4. 【解密结果】获得s3明文plaintext

注意：公钥私钥都输出为base64格式字符串保存，非pem格式。
AES CBC 256
RSA PKCS8 1024
*/

// AES+RSA 加密报文
func EnAesRsa(plaintext, pubKey string) (ciphertext string, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
		if err != nil {
			err = errors.New("EN_AES_RSA: 加密出错 " + err.Error())
		}
	}()
	if plaintext == "" {
		return
	}
	if pubKey == "" {
		err = errors.New("公钥不能为空")
		return
	}

	// 随机32位字符串作为AES密钥
	s1 := Rand32()
	aesKey := []byte(s1)

	// RSA  公钥支持pem格式与base格式
	var derBytes []byte
	if strings.Contains(pubKey, "BEGIN PUBLIC KEY") {
		p, _ := pem.Decode([]byte(pubKey))
		if p == nil {
			err = errors.New("pem.Decode error ")
			return
		}
		derBytes = p.Bytes
	} else {
		derBytes, err = base64.StdEncoding.DecodeString(pubKey)
		if err != nil {
			err = errors.New("base64.Decode error " + err.Error())
			return
		}
	}

	pub, err := x509.ParsePKIXPublicKey(derBytes)
	if err != nil {
		err = errors.New("x509.ParsePKIXPublicKey error " + err.Error())
		return
	}
	aes_key_ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, pub.(*rsa.PublicKey), aesKey)
	if err != nil {
		err = errors.New("rsa.EncryptPKCS1v15 error " + err.Error())
		return
	}
	s2 := base64.StdEncoding.EncodeToString(aes_key_ciphertext)

	// AES CBC 加密
	block, err := aes.NewCipher(aesKey)
	if err != nil {
		err = errors.New("aes.NewCipher error " + err.Error())
		return
	}
	plainText := Padding([]byte(plaintext), block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, []byte(s1[0:block.BlockSize()]))
	cipherText := make([]byte, len(plainText))
	blockMode.CryptBlocks(cipherText, plainText)
	s3 := base64.StdEncoding.EncodeToString(cipherText)

	// 密文结果
	ciphertext = s2 + ":" + s3
	return
}

// AES+RSA 解密报文
func DeAesRsa(ciphertext, priKey string) (plaintext string, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
		if err != nil {
			err = errors.New("DE_AES_RSA:解密出错" + err.Error())
		}
	}()
	if ciphertext == "" {
		return
	}
	if priKey == "" {
		err = errors.New("私钥不能为空")
		return
	}

	arr := strings.Split(ciphertext, ":")
	s2 := arr[0]
	s3 := arr[1]

	aes_key_ciphertext, err := base64.StdEncoding.DecodeString(s2)
	if err != nil {
		err = errors.New("base64.Decode aesKey error" + err.Error())
		return
	}
	cipherText, err := base64.StdEncoding.DecodeString(s3)
	if err != nil {
		err = errors.New("base64.Decode ciphertext  error" + err.Error())
		return
	}

	// RSA  公钥支持pem格式与base格式
	var der []byte
	if strings.Contains(priKey, "BEGIN PRIVATE KEY") {
		p, _ := pem.Decode([]byte(priKey))
		if p == nil {
			err = errors.New("pem.Decode privateKey error ")
			return
		}
		der = p.Bytes
	} else {
		der, err = base64.StdEncoding.DecodeString(priKey)
		if err != nil {
			err = errors.New("base64.Decode privateKey error" + err.Error())
			return
		}
	}

	prikey, err := x509.ParsePKCS8PrivateKey(der)
	if err != nil {
		err = errors.New("x509.ParsePKCS8PrivateKey error " + err.Error())
		return
	}
	key, err := rsa.DecryptPKCS1v15(rand.Reader, prikey.(*rsa.PrivateKey), aes_key_ciphertext)
	if err != nil {
		err = errors.New("rsa.DecryptPKCS1v15 error " + err.Error())
		return
	}

	// AES 解密
	block, err := aes.NewCipher(key)
	if err != nil {
		err = errors.New("aes.NewCipher error" + err.Error())
		return
	}
	blockMode := cipher.NewCBCDecrypter(block, []byte(string(key)[0:block.BlockSize()]))
	plainText := make([]byte, len(cipherText))
	blockMode.CryptBlocks(plainText, cipherText)
	plainText = UnPadding(plainText)

	// 明文结果
	plaintext = string(plainText)
	return
}

func Padding(plainText []byte, blockSize int) []byte {
	n := blockSize - len(plainText)%blockSize //计算要填充的长度
	temp := bytes.Repeat([]byte{byte(n)}, n)  //对原来的明文填充n个n
	plainText = append(plainText, temp...)
	return plainText
}

func UnPadding(cipherText []byte) []byte {
	end := cipherText[len(cipherText)-1]               //取出密文最后一个字节end
	cipherText = cipherText[:len(cipherText)-int(end)] //删除填充
	return cipherText
}

// Rand32 生成32位随机字符串 使用crypto/rand 随机赋值byte数组， 然后md5返回32位十六进制字符串
func Rand32() string {
	var b = make([]byte, 48)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return fmt.Sprintf("%x", md5.Sum(b))
}
