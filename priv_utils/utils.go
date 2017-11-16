package priv_utils

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"log"
	"os"
)

// 检查文件是否存在
func CheckFileIsExist(fileName string) (bool) {
	log.Println("utils, CheckFileIsExist", fileName)
	exist := true
	if _, err := os.Stat(fileName);os.IsNotExist(err) {
		exist = false
	}
	return exist
}

// // rsa加解密
// func RsaEncrypt(origData []byte) ([]byte, error) {
// 	//加密
// 	block, _ := pem.Decode(publicKey)
// 	if block == nil {
// 		return nil, errors.New("public key error")
// 	}
// 	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
// 	if err != nil {
// 		return nil, err
// 	}
// 	pub := pubInterface.(*rsa.PublicKey)
// 	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
// }

//DES 加解密
func DesEncrypt(origData, key[]byte) ([]byte,error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	origData = PKCS5Padding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key)
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func DesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, key)
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	return origData, nil
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length-unpadding)]
}