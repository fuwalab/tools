package util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

var hashed = md5.Sum([]byte("fuwalab"))
var commonIV = hashed[:]

// cipher key
var keyString = "tqax/vkLME-CjEP3##u/rVLi),XNGrZ-"

// Encrypt encrypt string
func Encrypt(s string) string {
	plaintext := []byte(s)

	c, err := aes.NewCipher([]byte(keyString))
	if err != nil {
		fmt.Printf("Error: NewCipher(%d bytes) = %s", len(keyString), err)
		panic(err)
	}

	cfb := cipher.NewCFBEncrypter(c, commonIV)
	cipherText := make([]byte, len(plaintext))
	cfb.XORKeyStream(cipherText, plaintext)

	return fmt.Sprintf("%x", cipherText)
}

// Decrypt decrypt string
func Decrypt(s string) string {
	c, err := aes.NewCipher([]byte(keyString))
	if err != nil {
		fmt.Printf("Error: NewCipher(%d bytes) = %s", len(keyString), err)
		panic(err)
	}

	ciphertext, _ := hex.DecodeString(s)
	cfbdec := cipher.NewCFBDecrypter(c, commonIV)
	plaintextCopy := make([]byte, len(ciphertext))
	cfbdec.XORKeyStream(plaintextCopy, ciphertext)

	return string(plaintextCopy)
}
