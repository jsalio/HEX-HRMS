package contracts

import "hrms.local/core/models"

type RoleContract interface {
	WriteOperation[models.Role]
	ReadOperation[models.Role]
	GetPermissions(roleID string) ([]models.Permission, *models.SystemError)
}
