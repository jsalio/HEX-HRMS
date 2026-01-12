package department

import (
	"hrms.local/core/contracts"
	"hrms.local/core/models"
)

type ListDepartmentUseCase struct {
	request    contracts.IGenericRequest[models.SearchQuery]
	repository contracts.DepartmentContract
}

func NewListDepartmentUseCase(request contracts.IGenericRequest[models.SearchQuery], repository contracts.DepartmentContract) *ListDepartmentUseCase {
	return &ListDepartmentUseCase{
		request:    request,
		repository: repository,
	}
}

func (u *ListDepartmentUseCase) Validate() *models.SystemError {
	request := u.request.Build()
	if err := request.Filters.Validate(models.Department{}); err != nil {
		return err
	}
	return nil
}

func (u *ListDepartmentUseCase) Execute() (*models.PaginatedResponse[models.Department], *models.SystemError) {
	request := u.request.Build()
	result, err := u.repository.GetByFilter(request)
	if err != nil {
		return nil, models.NewSystemError(models.SystemErrorCodeInternal, models.SystemErrorTypeValidation, models.SystemErrorLevelWarning, "Name already exists", nil)
	}
	return result, nil
}
