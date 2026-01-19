package position

import (
	"hrms.local/core/contracts"
	"hrms.local/core/models"
)

// DeletePositionUseCase handles deleting a position
type DeletePositionUseCase struct {
	request    contracts.IGenericRequest[models.Filter]
	repository contracts.PositionContract
}

// NewDeletePositionUseCase creates a new instance of DeletePositionUseCase
func NewDeletePositionUseCase(repository contracts.PositionContract, request contracts.IGenericRequest[models.Filter]) *DeletePositionUseCase {
	return &DeletePositionUseCase{
		request:    request,
		repository: repository,
	}
}

// Validate validates the request data
func (u *DeletePositionUseCase) Validate() *models.SystemError {
	request := u.request.Build()

	if request.Value == "" {
		return models.NewSystemError(models.SystemErrorCodeValidation, models.SystemErrorTypeValidation, models.SystemErrorLevelWarning, "Position ID is required", nil)
	}

	// Check if position exists
	_, err := u.repository.GetOnce("id", request.Value)
	if err != nil {
		return models.NewSystemError(models.SystemErrorCodeValidation, models.SystemErrorTypeValidation, models.SystemErrorLevelWarning, "Position not found", nil)
	}

	return nil
}

// Execute deletes the position from the repository
func (u *DeletePositionUseCase) Execute() (interface{}, *models.SystemError) {
	request := u.request.Build()

	_, err := u.repository.Delete(request.Value.(string))
	if err != nil {
		return nil, models.NewSystemError(models.SystemErrorCodeInternal, models.SystemErrorTypeInternal, models.SystemErrorLevelError, "Failed to delete position", nil)
	}

	return nil, nil
}
