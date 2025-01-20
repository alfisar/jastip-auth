package helper

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"jastip/config"
	"jastip/domain"
	"jastip/internal/errorhandler"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// func GeneratePass for generate password bcrypt
func GeneratePass(data string) (passHashed string, err domain.ErrorData) {
	passHash, errData := bcrypt.GenerateFromPassword([]byte(data), bcrypt.MinCost)

	if errData != nil {
		err = errorhandler.ErrHashing(errData)
		return
	}

	passHashed = string(passHash)
	return
}

// func Verify for verify password
func Verify(passwordHash string, password string) (err error) {
	err = bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	return
}

func TimeGenerator() string {
	timeNow := time.Now().Format(time.RFC3339)
	return timeNow
}

// Hash data string single line with method sha256
func HashSha256(data string) (result []byte) {
	keysHash := sha256.New()
	keysHash.Write([]byte(data))
	result = keysHash.Sum(nil)

	return
}

// Decode data with method base64
func Decode(s string) []byte {
	data, _ := base64.StdEncoding.DecodeString(s)
	return data
}

// Decode data with AES-256-CBC Method
func DecryptAES256CBC(poolData *config.Config, data string) (result string, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf(fmt.Sprintf("%s", r))
			return
		}
	}()

	key := HashSha256(poolData.Hash.Key) // 32 bytes for AES-256

	encryptedData := data
	encryptedDataByte := Decode(encryptedData)
	if len(encryptedDataByte) < aes.BlockSize {
		return "", fmt.Errorf("ciphertext too short")
	}

	iv := encryptedDataByte[:aes.BlockSize]
	encryptedDataByte = encryptedDataByte[aes.BlockSize:]
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	if len(encryptedDataByte) < aes.BlockSize {
		return "", fmt.Errorf("encryptedDataByte too short")
	}

	mode := cipher.NewCBCDecrypter(block, iv)

	mode.CryptBlocks(encryptedDataByte, encryptedDataByte)

	// Unpad the decrypted data
	return string(Unpad(encryptedDataByte)), nil
}

// Remove the padding data
func Unpad(data []byte) []byte {
	length := len(data)
	unpadding := int(data[length-1])
	return data[:(length - unpadding)]
}
