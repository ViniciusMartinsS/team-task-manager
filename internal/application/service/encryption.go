package service

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"

	"github.com/ViniciusMartinsS/manager/internal/domain"
)

type encryptionService struct{}

const keyString = "c286ff708a66d36d85bad3b4f00ae35b92a6f950671cf45375e8a9b007a11967" // env

func NewEncryption() domain.EncryptionService {
	return encryptionService{}
}

func (e encryptionService) Encrypt(content string) string {
	key, _ := hex.DecodeString(keyString)
	plaintext := []byte(content)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	ciphertext := aesGCM.Seal(nonce, nonce, plaintext, nil)
	return fmt.Sprintf("%x", ciphertext)
}

func (e encryptionService) Decrypt(contentEncrypted string) string {
	if contentEncrypted == "" {
		return ""
	}

	key, _ := hex.DecodeString(keyString)
	enc, _ := hex.DecodeString(contentEncrypted)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	nonceSize := aesGCM.NonceSize()
	nonce, ciphertext := enc[:nonceSize], enc[nonceSize:]

	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}

	return fmt.Sprintf("%s", plaintext)
}
