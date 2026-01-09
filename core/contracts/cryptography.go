package contracts

import "hrms.local/core/models"

type CryptographyContract interface {
	EncodePassword(string) (string, *models.SystemError)
	ComparePassword(string, string) (bool, *models.SystemError)
}
