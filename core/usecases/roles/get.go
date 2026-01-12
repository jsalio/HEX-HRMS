package roles

import (
	"hrms.local/core/contracts"
	"hrms.local/core/models"
)

type GetRoleUsecase struct {
	repo   contracts.RoleContract
	roleID string
}

func NewGetRoleUsecase(repo contracts.RoleContract, roleID string) *GetRoleUsecase {
	return &GetRoleUsecase{repo: repo, roleID: roleID}
}

func (u *GetRoleUsecase) Validate() *models.SystemError {
	if u.roleID == "" {
		return models.NewSystemError(models.SystemErrorCodeInternal, models.SystemErrorTypeValidation, models.SystemErrorLevelError, "Role ID is required", nil)
	}
	return nil
}

func (u *GetRoleUsecase) Execute() (*models.Role, *models.SystemError) {
	return u.repo.GetOnce("id", u.roleID)
}
