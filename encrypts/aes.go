package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"io"
	"log"
)

// pad adds PKCS#7 padding to a byte slice.
func pad(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padText...)
}

// unpad removes PKCS#7 padding from a byte slice.
func unpad(src []byte) ([]byte, error) {
	length := len(src)
	if length == 0 {
		return nil, errors.New("unpad error. Input length is 0")
	}
	padding := int(src[length-1])
	if padding < 1 || padding > aes.BlockSize {
		return nil, errors.New("unpad error. Invalid padding")
	}
	// Check padding
	padText := src[length-padding:]
	for _, v := range padText {
		if v != byte(padding) {
			return nil, errors.New("unpad error. Invalid padding")
		}
	}
	return src[:length-padding], nil
}

// encrypt encrypts plain text string using AES in CBC mode.
func encrypt(plainText string, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	plaintextBytes := pad([]byte(plainText), block.BlockSize())
	cipherText := make([]byte, aes.BlockSize+len(plaintextBytes))
	iv := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(cipherText[aes.BlockSize:], plaintextBytes)

	return hex.EncodeToString(cipherText), nil
}

// decrypt decrypts cipher text string using AES in CBC mode.
func decrypt(cipherText string, key []byte) (string, error) {
	cipherTextBytes, err := hex.DecodeString(cipherText)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	if len(cipherTextBytes) < aes.BlockSize {
		return "", errors.New("cipherText too short")
	}

	iv := cipherTextBytes[:aes.BlockSize]
	cipherTextBytes = cipherTextBytes[aes.BlockSize:]

	if len(cipherTextBytes)%aes.BlockSize != 0 {
		return "", errors.New("cipherText is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(cipherTextBytes, cipherTextBytes)

	plaintextBytes, err := unpad(cipherTextBytes)
	if err != nil {
		return "", err
	}

	return string(plaintextBytes), nil
}

func main() {
	key := []byte("thekeyhas32bytesthekeyhas32bytes") // AES-128, key should be 16, 24, or 32 bytes for AES-128, AES-192, or AES-256.
	plaintext := "Hello, AES Encryption in Go!"

	encrypted, err := encrypt(plaintext, key)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Encrypted: %s\n", encrypted)

	decrypted, err := decrypt(encrypted, key)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Decrypted: %s\n", decrypted)
}
