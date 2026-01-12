package roles

import (
	"hrms.local/core/contracts"
	"hrms.local/core/models"
)

type GetPermissionsUsecase struct {
	repo   contracts.RoleContract
	roleID string
}

func NewGetPermissionsUsecase(repo contracts.RoleContract, roleID string) *GetPermissionsUsecase {
	return &GetPermissionsUsecase{repo: repo, roleID: roleID}
}

func (u *GetPermissionsUsecase) Validate() *models.SystemError {
	if u.roleID == "" {
		return models.NewSystemError(models.SystemErrorCodeInternal, models.SystemErrorTypeValidation, models.SystemErrorLevelError, "Role ID is required", nil)
	}
	return nil
}

func (u *GetPermissionsUsecase) Execute() ([]models.Permission, *models.SystemError) {
	return u.repo.GetPermissions(u.roleID)
}
