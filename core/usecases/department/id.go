package department

import (
	"hrms.local/core/contracts"
	"hrms.local/core/models"
)

type GetDepartmentByIdUseCase struct {
	departmentContract contracts.DepartmentContract
	request            contracts.IGenericRequest[models.Filter]
}

func NewGetDepartmentByIdUseCase(departmentContract contracts.DepartmentContract, request contracts.IGenericRequest[models.Filter]) *GetDepartmentByIdUseCase {
	return &GetDepartmentByIdUseCase{
		departmentContract: departmentContract,
		request:            request,
	}
}

func (u *GetDepartmentByIdUseCase) Validate() *models.SystemError {
	request := u.request.Build()
	if request.Key == "" || request.Value == "" {
		return models.NewSystemError(models.SystemErrorCodeValidation, models.SystemErrorTypeValidation, models.SystemErrorLevelError, "Key and Value are required", nil)
	}
	if request.Key != "id" {
		return models.NewSystemError(models.SystemErrorCodeValidation, models.SystemErrorTypeValidation, models.SystemErrorLevelError, "Key must be id", nil)
	}
	exists, err := u.departmentContract.Exists(request.Key, request.Value)
	if err != nil {
		return models.NewSystemError(models.SystemErrorCodeInternal, models.SystemErrorTypeInternal, models.SystemErrorLevelError, "Error checking if user exists", nil)
	}
	if !exists {
		return models.NewSystemError(models.SystemErrorCodeNone, models.SystemErrorType(models.SystemErrorLevelError), models.SystemErrorLevelError, "User not found", nil)
	}

	return nil
}

func (u *GetDepartmentByIdUseCase) Execute() (*models.Department, *models.SystemError) {
	request := u.request.Build()
	department, err := u.departmentContract.GetOnce(request.Key, request.Value)
	if err != nil {
		return nil, models.NewSystemError(models.SystemErrorCodeInternal, models.SystemErrorTypeInternal, models.SystemErrorLevelError, "Error getting department", nil)
	}
	return department, nil
}
