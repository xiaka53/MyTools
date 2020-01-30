package encryption

import (
	"MyTools/src/utils"
	"crypto/aes"
	"encoding/hex"
)

//AES ECB模式加密
func EncAesEcb(Str string) {
	Key := getKey()
	cipher, _ := aes.NewCipher(generateKey(Key))
	length := (len(Str) + aes.BlockSize) / aes.BlockSize
	plain := make([]byte, length*aes.BlockSize)
	copy(plain, Str)
	pad := byte(len(plain) - len(Str))
	for i := len(Str); i < len(plain); i++ {
		plain[i] = pad
	}
	encrypted := make([]byte, len(plain))
	// 分组分块加密
	for bs, be := 0, cipher.BlockSize(); bs <= len(Str); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Encrypt(encrypted[bs:be], plain[bs:be])
	}
	Str = hex.EncodeToString(encrypted)
}

//AesEcb模式解密
func DecAesEcb() {
	cipher, _ := aes.NewCipher(generateKey(rncryprion.Key))
	encrypted, _ := hex.DecodeString(rncryprion.Str)
	decrypted := make([]byte, len(encrypted))
	for bs, be := 0, cipher.BlockSize(); bs < len(encrypted); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Decrypt(decrypted[bs:be], encrypted[bs:be])
	}
	trim := 0
	if len(decrypted) > 0 {
		trim = len(decrypted) - int(decrypted[len(decrypted)-1])
	}
	rncryprion.LowStr = decrypted[:trim]
}

//生成密钥
func generateKey(key []byte) (genKey []byte) {
	genKey = make([]byte, 16)
	copy(genKey, key)
	for i := 16; i < len(key); {
		for j := 0; j < 16 && i < len(key); j, i = j+1, i+1 {
			genKey[j] ^= key[i]
		}
	}
	return genKey
}

//获取密钥
func getKey() []byte {
	return []byte(utils.NumEngStr.Run(16))
}
