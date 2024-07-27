package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

func Encrypt(key, text []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	ciphertext := make([]byte, aes.BlockSize+len(text))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], text)

	return base64.URLEncoding.EncodeToString(ciphertext), nil
}

func Decrypt(key []byte, cryptoText string) (string, error) {
	ciphertext, _ := base64.URLEncoding.DecodeString(cryptoText)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	if len(ciphertext) < aes.BlockSize {
		return "", fmt.Errorf("ciphertext too short")
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)

	return string(ciphertext), nil
}

func main() {
	key := []byte("example key 1234") // key harus memiliki panjang 16, 24, atau 32 byte
	text := "Hello, World!"

	fmt.Println(text)
	encrypted, err := encrypt(key, []byte(text))
	if err != nil {
		fmt.Println("Error encrypting:", err)
		return
	}

	fmt.Printf("Encrypted: %s\n", encrypted)

	decrypted, err := decrypt(key, encrypted)
	if err != nil {
		fmt.Println("Error decrypting:", err)
		return
	}

	fmt.Printf("Decrypted: %s\n", decrypted)
}
