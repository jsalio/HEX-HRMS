package department

import (
	"hrms.local/core/contracts"
	"hrms.local/core/models"
)

type UpdateDepartmentUseCase struct {
	request    *contracts.GenericRequest[models.Department]
	repository contracts.DepartmentContract
}

func NewUpdateDepartmentUseCase(request *contracts.GenericRequest[models.Department], repository contracts.DepartmentContract) *UpdateDepartmentUseCase {
	return &UpdateDepartmentUseCase{
		request:    request,
		repository: repository,
	}
}

func (u *UpdateDepartmentUseCase) Validate() *models.SystemError {
	request := u.request.Build()

	if request.Name == "" {
		return models.NewSystemError(models.SystemErrorCodeInternal, models.SystemErrorTypeValidation, models.SystemErrorLevelWarning, "Name is empty", nil)
	}
	raw, err := u.repository.Exists("name", request.Name)
	if err != nil {
		return models.NewSystemError(models.SystemErrorCodeInternal, models.SystemErrorTypeValidation, models.SystemErrorLevelWarning, "Error on repository", nil)
	}
	if raw {
		return models.NewSystemError(models.SystemErrorCodeInternal, models.SystemErrorTypeValidation, models.SystemErrorLevelWarning, "Name already exists", nil)
	}
	return nil
}

func (u *UpdateDepartmentUseCase) Execute() (*models.Department, *models.SystemError) {
	request := u.request.Build()
	result, err := u.repository.Update(request.ID, request)
	if err != nil {
		return nil, models.NewSystemError(models.SystemErrorCodeInternal, models.SystemErrorTypeValidation, models.SystemErrorLevelWarning, "Name already exists", nil)
	}
	return &result, nil
}
