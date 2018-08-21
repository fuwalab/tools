package util

import "testing"

func TestEncrypt(t *testing.T) {
	plaintext := "test"
	expected := "225025e6"
	actual := Encrypt(plaintext)

	if actual != expected {
		t.Errorf("Not same values\nactual: %v, expected: %v\n", actual, expected)
	}
}

func TestDecrypt(t *testing.T) {
	cipherText := "225025e6"
	expected := "test"
	actual := Decrypt(cipherText)

	if actual != expected {
		t.Errorf("Not same values\nactual: %v, expected: %v\n", actual, expected)
	}
}
