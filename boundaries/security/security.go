package security

import (
	"hrms.local/core/contracts"
	"hrms.local/core/models"

	"golang.org/x/crypto/bcrypt"
)

type SecurityImpl struct {
}

func NewSecurityImpl() contracts.CryptographyContract {
	return &SecurityImpl{}
}

func (s *SecurityImpl) EncodePassword(password string) (string, *models.SystemError) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", models.NewSystemError(
			models.SystemErrorCodeInternal,
			models.SystemErrorTypeInternal,
			models.SystemErrorLevelError,
			"could not encrypt password: "+err.Error(),
			struct{}{},
		)
	}
	return string(bytes), nil
}

func (s *SecurityImpl) ComparePassword(password string, encodedPassword string) (bool, *models.SystemError) {
	err := bcrypt.CompareHashAndPassword([]byte(encodedPassword), []byte(password))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return false, nil
		}
		return false, models.NewSystemError(
			models.SystemErrorCodeInternal,
			models.SystemErrorTypeInternal,
			models.SystemErrorLevelError,
			"could not compare passwords: "+err.Error(),
			struct{}{},
		)
	}
	return true, nil
}
