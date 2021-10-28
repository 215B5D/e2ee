package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
)

func Encrypt(key string, data []byte) (string, error) {
	block, err := aes.NewCipher([]byte(key))

	if err != nil {
		return "", errors.New("unable to encrypt")
	}

	gcm, err := cipher.NewGCM(block)

	if err != nil {
		return "", errors.New("unable to encrypt")
	}

	nonce := make([]byte, gcm.NonceSize())

	if _, err = rand.Read(nonce); err != nil {
		return "", errors.New("unable to encrypt")
	}

	return hex.EncodeToString(gcm.Seal(nonce, nonce, data, nil)), nil
}

func Decrypt(key, data string) (string, error) {
	decoded, err := hex.DecodeString(data)

	if err != nil {
		return "", errors.New("unable to decrypt")
	}

	block, err := aes.NewCipher([]byte(key))

	if err != nil {
		return "", errors.New("unable to decrypt")
	}

	gcm, err := cipher.NewGCM(block)

	if err != nil {
		return "", errors.New("unable to decrypt")
	}

	nonce, str := decoded[:gcm.NonceSize()], decoded[gcm.NonceSize():]
	decrypted, err := gcm.Open(nil, nonce, str, nil)

	if err != nil {
		return "", errors.New("unable to decrypt")
	}

	return string(decrypted), nil
}
