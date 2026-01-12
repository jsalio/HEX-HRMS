package roles

import (
	"hrms.local/core/contracts"
	"hrms.local/core/models"
)

type CreateRoleUsecase struct {
	repo    contracts.RoleContract
	request *contracts.GenericRequest[models.CreateRole]
}

func NewCreateRoleUsecase(repo contracts.RoleContract, request *contracts.GenericRequest[models.CreateRole]) *CreateRoleUsecase {
	return &CreateRoleUsecase{repo: repo, request: request}
}

func (u *CreateRoleUsecase) Validate() *models.SystemError {
	request := u.request.Build()
	if request.Name == "" {
		return models.NewSystemError(models.SystemErrorCodeInternal, models.SystemErrorTypeValidation, models.SystemErrorLevelError, "El usuario ya existe", struct{}{})
	}
	if request.Description == "" {
		return models.NewSystemError(models.SystemErrorCodeInternal, models.SystemErrorTypeValidation, models.SystemErrorLevelError, "Description is required", nil)
	}
	if len(request.Permissions) == 0 {
		return models.NewSystemError(models.SystemErrorCodeInternal, models.SystemErrorTypeValidation, models.SystemErrorLevelError, "Permissions are required", nil)
	}
	return nil
}

func (u *CreateRoleUsecase) Execute() (*models.RoleItem, *models.SystemError) {
	request := u.request.Build()
	role := models.Role{
		Name:        request.Name,
		Description: request.Description,
		Permissions: request.Permissions,
	}
	createdRole, err := u.repo.Create(role)
	if err != nil {
		return nil, err
	}
	return createdRole.ToRoleItem(), nil
}
