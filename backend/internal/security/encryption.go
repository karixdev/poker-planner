package security

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

type IdEncryptor interface {
	Encrypt(id string) (string, error)
	Decrypt(encryptedId string) (string, error)
}

type SymmetricIdEncryptorConfig struct {
	Key string
}
type SymmetricIdEncryptor struct {
	gcm cipher.AEAD
}

func NewSymmetricIdEncryptor(config SymmetricIdEncryptorConfig) (*SymmetricIdEncryptor, error) {
	block, err := aes.NewCipher([]byte(config.Key))
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	return &SymmetricIdEncryptor{
		gcm: gcm,
	}, nil
}

func (e *SymmetricIdEncryptor) Encrypt(id string) (string, error) {
	nonce := make([]byte, e.gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	encryptedId := e.gcm.Seal(nonce, nonce, []byte(id), nil)
	return string(encryptedId), nil
}

func (e *SymmetricIdEncryptor) Decrypt(encryptedId string) (string, error) {
	nonceSize := e.gcm.NonceSize()
	if len(encryptedId) < nonceSize {
		return "", ErrInvalidEncryptedId
	}
	nonce, encryptedId := encryptedId[:nonceSize], encryptedId[nonceSize:]

	id, err := e.gcm.Open(nil, []byte(nonce), []byte(encryptedId), nil)
	if err != nil {
		return "", err
	}

	return string(id), nil
}
