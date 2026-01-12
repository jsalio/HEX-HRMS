package permissions

import (
	"hrms.local/core/contracts"
	"hrms.local/core/models"
)

type ListPermissionsUseCase struct {
	permissionContract contracts.PermissionContract
}

func NewListPermissionsUseCase(permissionContract contracts.PermissionContract) *ListPermissionsUseCase {
	return &ListPermissionsUseCase{
		permissionContract: permissionContract,
	}
}

func (u *ListPermissionsUseCase) Execute() ([]models.Permission, *models.SystemError) {
	return u.permissionContract.GetAll()
}
