package position

import (
	"hrms.local/core/contracts"
	"hrms.local/core/models"
)

// GetPositionByIdUseCase handles retrieving a position by ID
type GetPositionByIdUseCase struct {
	request    contracts.IGenericRequest[models.Filter]
	repository contracts.PositionContract
}

// NewGetPositionByIdUseCase creates a new instance of GetPositionByIdUseCase
func NewGetPositionByIdUseCase(repository contracts.PositionContract, request contracts.IGenericRequest[models.Filter]) *GetPositionByIdUseCase {
	return &GetPositionByIdUseCase{
		request:    request,
		repository: repository,
	}
}

// Validate validates the request data
func (u *GetPositionByIdUseCase) Validate() *models.SystemError {
	request := u.request.Build()

	if request.Value == "" {
		return models.NewSystemError(models.SystemErrorCodeValidation, models.SystemErrorTypeValidation, models.SystemErrorLevelWarning, "Position ID is required", nil)
	}

	return nil
}

// Execute retrieves the position from the repository
func (u *GetPositionByIdUseCase) Execute() (*models.Position, *models.SystemError) {
	request := u.request.Build()

	result, err := u.repository.GetOnce(request.Key, request.Value)
	if err != nil {
		return nil, models.NewSystemError(models.SystemErrorCodeValidation, models.SystemErrorTypeValidation, models.SystemErrorLevelWarning, "Position not found", nil)
	}

	return result, nil
}
