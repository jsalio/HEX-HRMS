package roles

import (
	"hrms.local/core/contracts"
	"hrms.local/core/models"
)

type DeleteRoleUsecase struct {
	repo   contracts.RoleContract
	roleID string
}

func NewDeleteRoleUsecase(repo contracts.RoleContract, roleID string) *DeleteRoleUsecase {
	return &DeleteRoleUsecase{repo: repo, roleID: roleID}
}

func (u *DeleteRoleUsecase) Validate() *models.SystemError {
	if u.roleID == "" {
		return models.NewSystemError(models.SystemErrorCodeInternal, models.SystemErrorTypeValidation, models.SystemErrorLevelError, "Role ID is required", nil)
	}
	return nil
}

func (u *DeleteRoleUsecase) Execute() *models.SystemError {
	_, err := u.repo.Delete(u.roleID)
	if err != nil {
		return models.NewSystemError(models.SystemErrorCodeInternal, models.SystemErrorTypeInternal, models.SystemErrorLevelError, err.Error(), nil)
	}
	return nil
}
