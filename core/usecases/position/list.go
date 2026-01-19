package position

import (
	"hrms.local/core/contracts"
	"hrms.local/core/models"
)

// ListPositionUseCase handles listing positions with pagination and filters
type ListPositionUseCase struct {
	request    contracts.IGenericRequest[models.SearchQuery]
	repository contracts.PositionContract
}

// NewListPositionUseCase creates a new instance of ListPositionUseCase
func NewListPositionUseCase(request contracts.IGenericRequest[models.SearchQuery], repository contracts.PositionContract) *ListPositionUseCase {
	return &ListPositionUseCase{
		request:    request,
		repository: repository,
	}
}

// Validate validates the search query filters
func (u *ListPositionUseCase) Validate() *models.SystemError {
	request := u.request.Build()
	if err := request.Filters.Validate(models.Position{}); err != nil {
		return err
	}
	return nil
}

// Execute retrieves positions based on the search query
func (u *ListPositionUseCase) Execute() (*models.PaginatedResponse[models.Position], *models.SystemError) {
	request := u.request.Build()
	result, err := u.repository.GetByFilter(request)
	if err != nil {
		return nil, models.NewSystemError(models.SystemErrorCodeInternal, models.SystemErrorTypeInternal, models.SystemErrorLevelError, "Failed to list positions", nil)
	}
	return result, nil
}
