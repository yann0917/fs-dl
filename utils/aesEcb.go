package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
	"errors"
)

var (
	ErrCipherKey           = errors.New("the secret key is wrong and cannot be decrypted. Please check")
	ErrKeyLengthSixteen    = errors.New("a sixteen or twenty-four or thirty-two length secret key is required")
	ErrKeyLengthEight      = errors.New("a eight-length secret key is required")
	ErrKeyLengthTwentyFour = errors.New("a twenty-four-length secret key is required")
	ErrPaddingSize         = errors.New("padding size error please check the secret key or iv")
	ErrIvAes               = errors.New("a sixteen-length ivaes is required")
	ErrIvDes               = errors.New("a eight-length ivdes key is required")
	ErrRsaBits             = errors.New("bits 1024 or 2048")
)

/*
Ecb is not recommended,use cbc
*/
type aesEcb struct {
	b         cipher.Block
	blockSize int
}

func newECB(b cipher.Block) *aesEcb {
	return &aesEcb{
		b:         b,
		blockSize: b.BlockSize(),
	}
}

type eachEncrypt aesEcb

func newECBEncrypter(b cipher.Block) cipher.BlockMode {
	return (*eachEncrypt)(newECB(b))
}

func (x *eachEncrypt) BlockSize() int { return x.blockSize }

func (x *eachEncrypt) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		return
	}
	if len(dst) < len(src) {
		return
	}

	for len(src) > 0 {
		x.b.Encrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}

type ebbDecrypt aesEcb

func newECBDecrypter(b cipher.Block) cipher.BlockMode {
	return (*ebbDecrypt)(newECB(b))
}

func (x *ebbDecrypt) BlockSize() int {
	return x.blockSize
}

func (x *ebbDecrypt) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		return
	}
	if len(dst) < len(src) {
		return
	}

	for len(src) > 0 {
		x.b.Decrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}

func AesEcbEncrypt(plainText, secretKey []byte) (cipherText []byte, err error) {
	if len(secretKey) != 16 && len(secretKey) != 24 && len(secretKey) != 32 {
		return nil, ErrKeyLengthSixteen
	}
	block, err := aes.NewCipher(secretKey)
	if err != nil {
		return nil, err
	}

	paddingText := PKCS5Padding(plainText, block.BlockSize())

	crypted := make([]byte, len(paddingText))
	encrypter := newECBEncrypter(block)
	encrypter.CryptBlocks(crypted, paddingText)

	return crypted, nil
}

func AesEcbDecrypt(plainText, secretKey []byte) (cipherText []byte, err error) {
	if len(secretKey) != 16 && len(secretKey) != 24 && len(secretKey) != 32 {
		return nil, ErrKeyLengthSixteen
	}
	block, err := aes.NewCipher(secretKey)
	if err != nil {
		return nil, err
	}

	ecbDecrypter := newECBDecrypter(block)
	decrypted := make([]byte, len(plainText))
	ecbDecrypter.CryptBlocks(decrypted, plainText)

	return PKCS5UnPadding(decrypted, ecbDecrypter.BlockSize())
}

func AesEcbEncryptBase64(plainText, key []byte) (cipherTextBase64 string, err error) {
	encryBytes, err := AesEcbEncrypt(plainText, key)
	return base64.StdEncoding.EncodeToString(encryBytes), err
}

func AesEcbEncryptHex(plainText, key []byte) (cipherTextHex string, err error) {
	encryBytes, err := AesEcbEncrypt(plainText, key)
	return hex.EncodeToString(encryBytes), err
}

func AesEcbDecryptByBase64(cipherTextBase64 string, key []byte) (plainText []byte, err error) {
	plainTextBytes, err := base64.StdEncoding.DecodeString(cipherTextBase64)
	if err != nil {
		return []byte{}, err
	}
	return AesEcbDecrypt(plainTextBytes, key)
}

func AesEcbDecryptByHex(cipherTextHex string, key []byte) (plainText []byte, err error) {
	plainTextBytes, err := hex.DecodeString(cipherTextHex)
	if err != nil {
		return []byte{}, err
	}
	return AesEcbDecrypt(plainTextBytes, key)
}

// It is populated using pkcs5

func PKCS5Padding(plainText []byte, blockSize int) []byte {
	padding := blockSize - (len(plainText) % blockSize)
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	newText := append(plainText, padText...)
	return newText
}

func PKCS5UnPadding(plainText []byte, blockSize int) ([]byte, error) {
	length := len(plainText)
	number := int(plainText[length-1])
	if number >= length || number > blockSize {
		return nil, ErrPaddingSize
	}
	return plainText[:length-number], nil
}
