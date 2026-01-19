package position

import (
	"hrms.local/core/contracts"
	"hrms.local/core/models"
)

// CreatePositionUseCase handles the creation of a new position
type CreatePositionUseCase struct {
	request    contracts.IGenericRequest[models.CreatePosition]
	repository contracts.PositionContract
}

// NewCreatePositionUseCase creates a new instance of CreatePositionUseCase
func NewCreatePositionUseCase(request contracts.IGenericRequest[models.CreatePosition], repository contracts.PositionContract) *CreatePositionUseCase {
	return &CreatePositionUseCase{
		request:    request,
		repository: repository,
	}
}

// Validate validates the request data
func (u *CreatePositionUseCase) Validate() *models.SystemError {
	request := u.request.Build()

	if request.Title == "" {
		return models.NewSystemError(models.SystemErrorCodeValidation, models.SystemErrorTypeValidation, models.SystemErrorLevelWarning, "Title is required", nil)
	}
	if request.Code == "" {
		return models.NewSystemError(models.SystemErrorCodeValidation, models.SystemErrorTypeValidation, models.SystemErrorLevelWarning, "Code is required", nil)
	}
	if request.DepartmentID == "" {
		return models.NewSystemError(models.SystemErrorCodeValidation, models.SystemErrorTypeValidation, models.SystemErrorLevelWarning, "DepartmentID is required", nil)
	}
	if request.MaxEmployees <= 0 {
		return models.NewSystemError(models.SystemErrorCodeValidation, models.SystemErrorTypeValidation, models.SystemErrorLevelWarning, "MaxEmployees must be greater than 0", nil)
	}

	// Check if position code already exists
	exists, _ := u.repository.Exists("code", request.Code)
	if exists {
		return models.NewSystemError(models.SystemErrorCodeValidation, models.SystemErrorTypeValidation, models.SystemErrorLevelWarning, "Position code already exists", nil)
	}

	return nil
}

// Execute creates the position in the repository
func (u *CreatePositionUseCase) Execute() (*models.Position, *models.SystemError) {
	request := u.request.Build()
	position := request.ToPosition()

	result, err := u.repository.Create(*position)
	if err != nil {
		return nil, models.NewSystemError(models.SystemErrorCodeInternal, models.SystemErrorTypeInternal, models.SystemErrorLevelError, "Failed to create position", nil)
	}

	return &result, nil
}
