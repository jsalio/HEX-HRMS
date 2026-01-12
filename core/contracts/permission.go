package contracts

import "hrms.local/core/models"

type PermissionContract interface {
	GetAll() ([]models.Permission, *models.SystemError)
}
