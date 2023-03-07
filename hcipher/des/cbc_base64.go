package des

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"encoding/base64"
)

/* DES CBC PKCS5Padding  */

// EncodeCBCBase64	DES加密 CBC
func EncodeCBCBase64(s, key string) (v string, err error) {
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
	v = base64.StdEncoding.EncodeToString(cipherText)
	return
}

// DecodeCBCBase64	DES解密 CBC
func DecodeCBCBase64(s, key string) (v string, err error) {
	// Base64 Decode
	b1, err := base64.StdEncoding.DecodeString(s)
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

	v = string(text)
	return
}
