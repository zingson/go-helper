package metro

import (
	"crypto/aes"
)

func AesEncryptECB(origData []byte, key []byte) (encrypted []byte) {
	// key长度只能32位
	cipher, _ := aes.NewCipher(generateKey(key, 32))
	length := (len(origData) + aes.BlockSize) / aes.BlockSize
	plain := make([]byte, length*aes.BlockSize)
	copy(plain, origData)
	pad := byte(len(plain) - len(origData))
	for i := len(origData); i < len(plain); i++ {
		plain[i] = pad
	}
	encrypted = make([]byte, len(plain))
	// 分组分块加密
	for bs, be := 0, cipher.BlockSize(); bs <= len(origData); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Encrypt(encrypted[bs:be], plain[bs:be])
	}
	return encrypted
}
func AesDecryptECB(encrypted []byte, key []byte) (decrypted []byte) {
	cipher, _ := aes.NewCipher(generateKey(key, 32))
	decrypted = make([]byte, len(encrypted))
	//
	for bs, be := 0, cipher.BlockSize(); bs < len(encrypted); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Decrypt(decrypted[bs:be], encrypted[bs:be])
	}

	trim := 0
	if len(decrypted) > 0 {
		trim = len(decrypted) - int(decrypted[len(decrypted)-1])
	}
	return decrypted[:trim]
}
func generateKey(key []byte, klen int) (genKey []byte) {
	genKey = make([]byte, klen)
	copy(genKey, key)
	for i := klen; i < len(key); {
		for j := 0; j < klen && i < len(key); j, i = j+1, i+1 {
			genKey[j] ^= key[i]
		}
	}
	return genKey
}
