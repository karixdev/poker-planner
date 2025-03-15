package security

import "testing"

func TestIfReturnsErrorWhenInvalidKeyIsGiven(t *testing.T) {
	// Given
	config := SymmetricIdEncryptorConfig{
		Key: "123",
	}

	// When
	encryptor, err := NewSymmetricIdEncryptor(config)

	// Then
	if encryptor != nil {
		t.Fatalf("encryptor is not nil")
	}
	if err == nil {
		t.Errorf("err is nil")
	}
}

func TestIfReturnsEncryptorWhenKeyIsValid(t *testing.T) {
	// Given
	config := SymmetricIdEncryptorConfig{
		Key: "2b7e151628aed2a6abf7158809cf4f3c",
	}

	// When
	encryptor, err := NewSymmetricIdEncryptor(config)

	// Then
	if err != nil {
		t.Fatalf("err is not nil: %v", err)
	}
	if encryptor == nil {
		t.Errorf("encryptor is nil")
	}
}

func TestIfReturnsErrInvalidEncryptedIdIfEncryptedIdIsTooShort(t *testing.T) {
	// Given
	encryptor, _ := NewSymmetricIdEncryptor(SymmetricIdEncryptorConfig{
		Key: "2b7e151628aed2a6abf7158809cf4f3c",
	})
	encrypted := ""
	for range encryptor.gcm.NonceSize() - 1 {
		encrypted = encrypted + "1"
	}

	// When
	_, err := encryptor.Decrypt(encrypted)

	// Then
	if err == nil {
		t.Error("err is nil")
	}
}

func TestIfReturnsErrorIfIdWasNotCorrectlyEncrypted(t *testing.T) {
	// Given
	encryptor, _ := NewSymmetricIdEncryptor(SymmetricIdEncryptorConfig{
		Key: "2b7e151628aed2a6abf7158809cf4f3c",
	})
	encrypted := ""
	for range encryptor.gcm.NonceSize() {
		encrypted = encrypted + "1"
	}

	// When
	_, err := encryptor.Decrypt(encrypted)

	// Then
	if err == nil {
		t.Error("err is nil")
	}
}

func TestIfCorrectlyEncryptsAndDecryptsId(t *testing.T) {
	// Given
	encryptor, _ := NewSymmetricIdEncryptor(SymmetricIdEncryptorConfig{
		Key: "2b7e151628aed2a6abf7158809cf4f3c",
	})
	id := "123"

	// When
	encrypted, err := encryptor.Encrypt(id)
	if err != nil {
		t.Fatalf("could not encrypt: %v", err)
	}

	decrypted, err := encryptor.Decrypt(encrypted)
	if err != nil {
		t.Fatalf("could not decrypt: %v", err)
	}

	// Then
	if id != decrypted {
		t.Errorf("decrypted id is not equal to source: %s", decrypted)
	}
}
