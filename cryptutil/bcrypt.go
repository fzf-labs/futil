package cryptutil

import (
	"github.com/golang-module/dongle"
	"golang.org/x/crypto/bcrypt"
)

// Encrypt encrypts the plain text with bcrypt.
func Encrypt(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedBytes), err
}

// Compare compares the encrypted text with the plain text if it's the same.
func Compare(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func BcryptSign(data string) string {
	return dongle.Sign.FromString(data).ByBcrypt().ToRawString()
}

func BcryptVerify(sign, data string) bool {
	return dongle.Verify.FromRawString(sign, data).ByBcrypt().ToBool()
}
