package v3

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"strings"
)

/**
V3 报文加密（AES）+ 签名（SHA1） +密钥加密（RSA）方式方法如下：

加密步骤：
1. 明文字符串 plaintext ;
2. 生成16位随机字符串 ak  作为AES与SHA1的秘钥 ;
3. AES-CBC加密 plaintext 输出base64字符串 s1 ;
4. SHA1计算字符串 s1+ak 签名，输出 s2 ;
5. RSA加密 ak 输出s3 ;
6. 获取最终密文 ciphertext = s1:s2:s3

解密步骤：
1. 密文字符串 ciphertext
2. 根据冒号拆分获取s1,s2,s3;
3. RSA 解密s3，获取 ak ，用于AES解密与SHA1验签；
4. SHA1 计算 s1+ak 签名 sign，
5. 验证签名sign是否等于s2；
6. AES解密s1获取明文 plaintext ;

注意：公钥私钥都输出为base64格式字符串保存，非pem格式。
AES CBC 128
RSA PKCS8 1024
*/

// Encode AES+SHA1+RSA 加密报文
func Encode(plaintext, pubKey string) (ciphertext string, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
		if err != nil {
			err = errors.New("ENCODE_V3_ERROR: 加密出错 " + err.Error())
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
	aKey := rand32()[8:24]
	aesKey := []byte(aKey)

	// AES CBC 128 加密
	block, err := aes.NewCipher(aesKey)
	if err != nil {
		err = errors.New("aes.NewCipher error " + err.Error())
		return
	}
	plainText := Padding([]byte(plaintext), block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, []byte(aKey[0:block.BlockSize()]))
	cipherText := make([]byte, len(plainText))
	blockMode.CryptBlocks(cipherText, plainText)
	s1 := base64.StdEncoding.EncodeToString(cipherText)

	// SHA1签名
	s2 := fmt.Sprintf("%x", sha1.Sum([]byte(s1+aKey)))

	// RSA  公钥支持pem格式与base格式
	var derBytes []byte
	if strings.Contains(pubKey, "BEGIN PUBLIC KEY") {
		p, _ := pem.Decode([]byte(pubKey))
		if p == nil {
			err = errors.New("pubKey pem.Decode error ")
			return
		}
		derBytes = p.Bytes
	} else {
		derBytes, err = base64.StdEncoding.DecodeString(pubKey)
		if err != nil {
			err = errors.New("pubKey Base64解码错误 " + err.Error())
			return
		}
	}

	pub, err := x509.ParsePKIXPublicKey(derBytes)
	if err != nil {
		err = errors.New("x509.ParsePKIXPublicKey error " + err.Error())
		return
	}
	aesKeyCiphertext, err := rsa.EncryptPKCS1v15(rand.Reader, pub.(*rsa.PublicKey), aesKey)
	if err != nil {
		err = errors.New("rsa.EncryptPKCS1v15 error " + err.Error())
		return
	}
	s3 := base64.StdEncoding.EncodeToString(aesKeyCiphertext)

	// 密文结果 =  报文密文:签名:密钥密文
	ciphertext = s1 + ":" + s2 + ":" + s3
	return
}

// Decode AES+SHA1+RSA 解密报文
func Decode(ciphertext, priKey string) (plaintext string, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
		if err != nil {
			err = errors.New("DECODE_V3_ERROR:" + err.Error())
		}
	}()
	if ciphertext == "" {
		return
	}
	if priKey == "" {
		err = errors.New("私钥不能为空")
		return
	}

	split := strings.Split(ciphertext, ":")
	s1 := split[0] // 密文报文
	s2 := split[1] // s1的签名，需SHA1验签
	s3 := split[2] // s1的秘钥，需RSA解密

	// RSA  公钥支持pem格式与base格式
	var der []byte
	if strings.Contains(priKey, "BEGIN PRIVATE KEY") {
		p, _ := pem.Decode([]byte(priKey))
		if p == nil {
			err = errors.New("priKey pem.Decode error ")
			return
		}
		der = p.Bytes
	} else {
		der, err = base64.StdEncoding.DecodeString(priKey)
		if err != nil {
			err = errors.New("priKey Base64解码错误 " + err.Error())
			return
		}
	}

	// RSA 解密
	aesKeyCiphertext, err := base64.StdEncoding.DecodeString(s3)
	if err != nil {
		err = errors.New("s3 Base64解码错误 " + err.Error())
		return
	}
	privateKey, err := x509.ParsePKCS8PrivateKey(der)
	if err != nil {
		err = errors.New("x509.ParsePKCS8PrivateKey error " + err.Error())
		return
	}
	aKeyBytes, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey.(*rsa.PrivateKey), aesKeyCiphertext)
	if err != nil {
		err = errors.New("rsa.DecryptPKCS1v15 error " + err.Error())
		return
	}
	aKey := string(aKeyBytes)

	// SHA1验签
	sign := fmt.Sprintf("%x", sha1.Sum([]byte(s1+aKey)))
	if sign != s2 {
		err = errors.New("验签失败")
		return
	}

	// AES 解密
	cipherText, err := base64.StdEncoding.DecodeString(s1)
	if err != nil {
		err = errors.New("s1 Base64解码错误 " + err.Error())
		return
	}
	block, err := aes.NewCipher(aKeyBytes)
	if err != nil {
		err = errors.New("aes.NewCipher error " + err.Error())
		return
	}
	blockMode := cipher.NewCBCDecrypter(block, []byte(aKey[0:block.BlockSize()]))
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
func rand32() string {
	var b = make([]byte, 48)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return fmt.Sprintf("%x", md5.Sum(b))
}
