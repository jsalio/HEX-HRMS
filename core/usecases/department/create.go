package department

import (
	"hrms.local/core/contracts"
	"hrms.local/core/models"
)

type CreateDepartmentUseCase struct {
	request    *contracts.GenericRequest[models.Department]
	repository contracts.DepartmentContract
}

func NewCreateDepartmentUseCase(request *contracts.GenericRequest[models.Department], repository contracts.DepartmentContract) *CreateDepartmentUseCase {
	return &CreateDepartmentUseCase{
		request:    request,
		repository: repository,
	}
}

func (u *CreateDepartmentUseCase) Validate() *models.SystemError {
	request := u.request.Build()

	if request.Name == "" {
		return models.NewSystemError(models.SystemErrorCodeInternal, models.SystemErrorTypeValidation, models.SystemErrorLevelWarning, "Name is empty", nil)
	}
	raw, _ := u.repository.Exists("name", request.Name)
	if raw {
		return models.NewSystemError(models.SystemErrorCodeInternal, models.SystemErrorTypeValidation, models.SystemErrorLevelWarning, "Name already exists", nil)
	}
	return nil
}

func (u *CreateDepartmentUseCase) Execute() (*models.Department, *models.SystemError) {
	request := u.request.Build()
	result, err := u.repository.Create(request)
	if err != nil {
		return nil, models.NewSystemError(models.SystemErrorCodeInternal, models.SystemErrorTypeValidation, models.SystemErrorLevelWarning, "Name already exists", nil)
	}
	return &result, nil
}
