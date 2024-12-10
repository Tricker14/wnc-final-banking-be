package passwordencoderimplement

import (
	"github.com/21CLC01-WNC-Banking/WNC-Banking-BE/internal/utils/password_encoder"
	"golang.org/x/crypto/bcrypt"
)

type BcryptPasswordEncoder struct{}

func NewBcryptPasswordEncoder() password_encoder.PasswordEncoder {
	return &BcryptPasswordEncoder{}
}

func (b *BcryptPasswordEncoder) Encrypt(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func (b *BcryptPasswordEncoder) Compare(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
