package position

import (
	"hrms.local/core/contracts"
	"hrms.local/core/models"
)

// UpdatePositionUseCase handles updating an existing position
type UpdatePositionUseCase struct {
	request    contracts.IGenericRequest[models.ModifyPosition]
	repository contracts.PositionContract
}

// NewUpdatePositionUseCase creates a new instance of UpdatePositionUseCase
func NewUpdatePositionUseCase(request contracts.IGenericRequest[models.ModifyPosition], repository contracts.PositionContract) *UpdatePositionUseCase {
	return &UpdatePositionUseCase{
		request:    request,
		repository: repository,
	}
}

// Validate validates the request data
func (u *UpdatePositionUseCase) Validate() *models.SystemError {
	request := u.request.Build()

	if request.ID == "" {
		return models.NewSystemError(models.SystemErrorCodeValidation, models.SystemErrorTypeValidation, models.SystemErrorLevelWarning, "ID is required", nil)
	}
	if request.Title == "" {
		return models.NewSystemError(models.SystemErrorCodeValidation, models.SystemErrorTypeValidation, models.SystemErrorLevelWarning, "Title is required", nil)
	}
	if request.Code == "" {
		return models.NewSystemError(models.SystemErrorCodeValidation, models.SystemErrorTypeValidation, models.SystemErrorLevelWarning, "Code is required", nil)
	}
	if request.DepartmentID == "" {
		return models.NewSystemError(models.SystemErrorCodeValidation, models.SystemErrorTypeValidation, models.SystemErrorLevelWarning, "DepartmentID is required", nil)
	}

	// Check if position exists
	_, err := u.repository.GetOnce("id", request.ID)
	if err != nil {
		return models.NewSystemError(models.SystemErrorCodeValidation, models.SystemErrorTypeValidation, models.SystemErrorLevelWarning, "Position not found", nil)
	}

	return nil
}

// Execute updates the position in the repository
func (u *UpdatePositionUseCase) Execute() (*models.Position, *models.SystemError) {
	request := u.request.Build()
	position := request.ToPosition()

	result, err := u.repository.Update(position.ID, *position)
	if err != nil {
		return nil, models.NewSystemError(models.SystemErrorCodeInternal, models.SystemErrorTypeInternal, models.SystemErrorLevelError, "Failed to update position", nil)
	}

	return &result, nil
}
