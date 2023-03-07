package metro2

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
)

//EncodeCBC  AES CBC 加密
func EncodeCBC(plaintext, key, iv string) (ciphertext string, err error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		err = errors.New("aes.NewCipher error " + err.Error())
		return
	}
	plainText := padding([]byte(plaintext), block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, []byte(iv))
	cipherText := make([]byte, len(plainText))
	blockMode.CryptBlocks(cipherText, plainText)
	ciphertext = base64.StdEncoding.EncodeToString(cipherText)
	return
}

//DecodeCBC  AES CBC 解密
func DecodeCBC(ciphertext, key, iv string) (plaintext string, err error) {
	cipherText, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		err = errors.New("base64.Decode ciphertext  error" + err.Error())
		return
	}

	// AES 解密
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		err = errors.New("aes.NewCipher error" + err.Error())
		return
	}
	blockMode := cipher.NewCBCDecrypter(block, []byte(iv))
	plainText := make([]byte, len(cipherText))
	blockMode.CryptBlocks(plainText, cipherText)
	plainText = unPadding(plainText)

	// 明文结果
	plaintext = string(plainText)
	return
}

func padding(plainText []byte, blockSize int) []byte {
	n := blockSize - len(plainText)%blockSize //计算要填充的长度
	temp := bytes.Repeat([]byte{byte(n)}, n)  //对原来的明文填充n个n
	plainText = append(plainText, temp...)
	return plainText
}

func unPadding(plainText []byte) []byte {
	end := plainText[len(plainText)-1]              //取出密文最后一个字节end
	plainText = plainText[:len(plainText)-int(end)] //删除填充
	return plainText
}
