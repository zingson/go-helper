package desmd5

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"crypto/md5"
	"encoding/base64"
	"errors"
	"fmt"
)

/*
对称加密（DES）+验签（MD5）方式方法如下：

加密步骤：
1. 【加密报文】通过双方约定的密钥K1使用DES加密报文明文S1，得到密文字节数组B1;
2. 【加密报文】取得密文字节数组的Base64表示形式S2;
3. 【计算签名】将双方约定的密钥K1拼接在S2的前部，得到S3;
4. 【计算签名】计算S3的MD5，并取得Hex字符串的小写形式M1;
5. 【拼接结果】将M1和S2拼接得到最终密文报文：R

解密步骤：
1. 【拆解密文】取得密文报文字符串S1的前32个字符，得到签名M1，取得32个字符后的内容，得到S2;
2. 【验证签名】将双方约定的密钥K1拼接在S2的前部，得到S3;
3. 【验证签名】计算S3的MD5，并取得Hex字符串的小写形式M2;
4. 【验证签名】判断M1是否等于M2，不相等则验签失败，否则取得S2的反Base64后的密文字节数组B1;
5. 【解密报文】通过双方约定的密钥K1使用DES解密密文字节数组B1，得到明文报文：R

*/

// DES+MD5接口报文加密
func EnDesMd5(s, key string) (v string, err error) {
	k1 := []byte(key)
	s1 := []byte(s)

	block, err := des.NewCipher(k1)
	if err != nil {
		return
	}
	// PKCS5Padding
	blockSize := block.BlockSize()
	padding := blockSize - len(s1)%blockSize
	paddingText := bytes.Repeat([]byte{byte(padding)}, padding)
	s1 = append(s1, paddingText...)

	//DES CBCEncrypt
	blockMode := cipher.NewCBCEncrypter(block, k1)
	cipherText := make([]byte, len(s1))
	blockMode.CryptBlocks(cipherText, s1)

	// Base64 Encode
	s2 := base64.StdEncoding.EncodeToString(cipherText)

	// Md5 签名
	s3 := key + s2
	m1 := fmt.Sprintf("%x", md5.Sum([]byte(s3)))
	r := m1 + s2
	return r, err
}

// DES+MD5接口报文解密
func DeDesMd5(s, key string) (v string, err error) {
	if len(s) < 32 {
		err = errors.New("Error:Ciphertext format error")
		return
	}
	m1 := s[0:32]
	s2 := s[32:]

	// MD5签名验证
	s3 := key + s2
	m2 := fmt.Sprintf("%x", md5.Sum([]byte(s3)))
	if m1 != m2 {
		err = errors.New("Error:Signature verification failed")
		return
	}

	// Base64 Decode
	b1, err := base64.StdEncoding.DecodeString(s2)
	if err != nil {
		return
	}

	k1 := []byte(key)
	block, err := des.NewCipher(k1)
	if err != nil {
		return
	}

	//DES CBCDecrypt
	blockMode := cipher.NewCBCDecrypter(block, k1)
	text := make([]byte, len(b1))
	blockMode.CryptBlocks(text, b1)

	// PKCS5UnPadding
	length := len(text)
	unPadding := int(text[length-1])
	text = text[:(length - unPadding)]

	r := string(text)
	return r, err
}
