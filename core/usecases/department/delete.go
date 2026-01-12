package department

import (
	"hrms.local/core/contracts"
	"hrms.local/core/models"
)

type DeleteDepartmentUsecase struct {
	repository contracts.DepartmentContract
	request    contracts.IGenericRequest[models.Filter]
}

func NewDeleteDepartmentUsecase(repository contracts.DepartmentContract, request contracts.IGenericRequest[models.Filter]) *DeleteDepartmentUsecase {
	return &DeleteDepartmentUsecase{
		repository: repository,
		request:    request,
	}
}

func (u *DeleteDepartmentUsecase) Validate() *models.SystemError {
	request := u.request.Build()
	if request.Key == "" || request.Value == "" {
		return models.NewSystemError(models.SystemErrorCodeValidation, models.SystemErrorTypeValidation, models.SystemErrorLevelError, "Key and Value are required", nil)
	}
	if request.Key != "id" {
		return models.NewSystemError(models.SystemErrorCodeValidation, models.SystemErrorTypeValidation, models.SystemErrorLevelError, "Key must be id", nil)
	}
	exists, err := u.repository.Exists(request.Key, request.Value)
	if err != nil {
		return models.NewSystemError(models.SystemErrorCodeInternal, models.SystemErrorTypeInternal, models.SystemErrorLevelError, "Error checking if user exists", nil)
	}
	if !exists {
		return models.NewSystemError(models.SystemErrorCodeNone, models.SystemErrorType(models.SystemErrorLevelError), models.SystemErrorLevelError, "User not found", nil)
	}

	return nil
}

func (u *DeleteDepartmentUsecase) Execute() (*interface{}, *models.SystemError) {
	request := u.request.Build()
	department, err := u.repository.Delete(request.Value.(string))
	if err != nil {
		return nil, models.NewSystemError(models.SystemErrorCodeInternal, models.SystemErrorTypeInternal, models.SystemErrorLevelError, "Error deleting department", nil)
	}
	return &department, nil
}
