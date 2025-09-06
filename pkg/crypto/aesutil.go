package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
	"os"
)

func getKey() ([]byte, error) {
	b64 := os.Getenv("AES_KEY_BASE64")
	if b64 == "" {
		return nil, errors.New("Key not set")
	}
	key, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		return nil, err
	}
	if len(key) != 32 {
		return nil, errors.New("Key length must be 32 bytes")
	}
	return key, nil
}

func Encrypt(textNya string) (string, error) {
	key, err := getKey()
	if err != nil {
		return "", err
	}
	
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	ciphertext := gcm.Seal(nonce, nonce, []byte(textNya), nil)
	combined := append(nonce, ciphertext...)
	return base64.StdEncoding.EncodeToString(combined), nil
}

func Decrypt(base64Ciphertext string) (string, error) {
	key, err := getKey()
	if err != nil { return "", err }
	data, err := base64.StdEncoding.DecodeString(base64Ciphertext)
	if err != nil { return "", err }
	block, err := aes.NewCipher(key)
	if err != nil { return "", err }
	gcm, err := cipher.NewGCM(block)
	if err != nil { return "", err }
	if len(data) < gcm.NonceSize() { return "", errors.New("ciphertext too short") }
	nonce := data[:gcm.NonceSize()]
	ct := data[gcm.NonceSize():]
	pt, err := gcm.Open(nil, nonce, ct, nil)
	if err != nil { return "", err }
	return string(pt), nil
}