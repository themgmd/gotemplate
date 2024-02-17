package cipher

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

const defaultKeySize = 16

type Cipher struct {
	key []byte
}

func New(key []byte) *Cipher {
	return &Cipher{key: key}
}

func (c Cipher) Encode(plaintext string) (string, error) {
	block, err := aes.NewCipher(c.key)
	if err != nil {
		return "", fmt.Errorf("aes.NewCipher: %w", err)
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", fmt.Errorf("cipher.NewGCM: %w", err)
	}

	nonce := make([]byte, aesgcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", fmt.Errorf("io.ReadFull: %w", err)
	}

	ciphertext := aesgcm.Seal(nonce, nonce, []byte(plaintext), nil)
	encoded := fmt.Sprintf("%x", ciphertext)
	return encoded, nil
}

func (c Cipher) Decode(encoded string) (string, error) {
	enc, _ := hex.DecodeString(encoded)

	block, err := aes.NewCipher(c.key)
	if err != nil {
		return "", fmt.Errorf("aes.NewCipher: %w", err)
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", fmt.Errorf("cipher.NewGCM: %w", err)
	}

	nonceSize := aesgcm.NonceSize()
	nonce, ciphertext := enc[:nonceSize], enc[nonceSize:]

	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", fmt.Errorf("aesgcm.Open: %w", err)
	}

	return fmt.Sprintf("%s", plaintext), nil
}

func NewKey() ([]byte, error) {
	key := make([]byte, defaultKeySize)
	if _, err := io.ReadFull(rand.Reader, key); err != nil {
		return nil, fmt.Errorf("io.ReadFull: %w", err)
	}

	return key, nil
}
