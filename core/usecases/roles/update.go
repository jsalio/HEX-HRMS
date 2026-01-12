package roles

import (
	"hrms.local/core/contracts"
	"hrms.local/core/models"
)

type UpdateRoleUsecase struct {
	repo    contracts.RoleContract
	request *contracts.GenericRequest[models.Role]
}

func NewUpdateRoleUsecase(repo contracts.RoleContract, request *contracts.GenericRequest[models.Role]) *UpdateRoleUsecase {
	return &UpdateRoleUsecase{repo: repo, request: request}
}

func (u *UpdateRoleUsecase) Validate() *models.SystemError {
	request := u.request.Build()
	if request.ID == "" {
		return models.NewSystemError(models.SystemErrorCodeInternal, models.SystemErrorTypeValidation, models.SystemErrorLevelError, "ID is required", nil)
	}
	if request.Name == "" {
		return models.NewSystemError(models.SystemErrorCodeInternal, models.SystemErrorTypeValidation, models.SystemErrorLevelError, "Name is required", nil)
	}
	return nil
}

func (u *UpdateRoleUsecase) Execute() (*models.RoleItem, *models.SystemError) {
	request := u.request.Build()
	updatedRole, err := u.repo.Update(request.ID, request)
	if err != nil {
		return nil, err
	}
	return updatedRole.ToRoleItem(), nil
}
