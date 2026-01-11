package user

import (
	"hrms.local/core/contracts"
	"hrms.local/core/models"
)

// ListUserUseCase handles the retrieval of users based on specific filters.
// It orchestrates the validation of filters and the fetching of data via the user contract.
//
// Example Usage:
//
//	// 1. Create the repository/contract implementation
//	var userRepo contracts.UserContract = ... // Your repository implementation
//
//	// 2. Create the request with filters
//	filters := []models.Filter{
//		{Key: "Username", Value: "jdoe"},
//		{Key: "Type", Value: models.UserTypeAdmin},
//	}
//	// Assuming you have a concrete implementation of IGenericRequest or use the struct directly if generic
//	request := &contracts.GenericRequest[[]models.Filter]{Data: filters}
//
//	// 3. Instantiate the UseCase
//	useCase := user.NewListUserUseCase(userRepo, request)
//
//	// 4. Validate the request
//	if err := useCase.Validate(); err != nil {
//	    log.Printf("Validation failed: %v", err.Message)
//	    return
//	}
//
//	// 5. Execute the logic
//	users, err := useCase.Execute()
//	if err != nil {
//	    log.Printf("Execution failed: %v", err.Message)
//	    return
//	}
//
//	for _, user := range users {
//	    fmt.Printf("User: %s\n", user.Username)
//	}
type ListUserUseCase struct {
	userContract contracts.UserContract
	request      contracts.IGenericRequest[models.SearchQuery]
}

// NewListUserUseCase creates a new instance of ListUserUseCase.
// It injects the user contract (dependency inversion) and the request data.
func NewListUserUseCase(userContract contracts.UserContract, request contracts.IGenericRequest[models.SearchQuery]) *ListUserUseCase {
	return &ListUserUseCase{
		userContract: userContract,
		request:      request,
	}
}

// Validate ensures that the request filters are valid.
// It checks if the request is empty and validates each filter against the User model using reflection.
func (u *ListUserUseCase) Validate() *models.SystemError {
	request := u.request.Build()
	if err := request.Filters.Validate(models.User{}); err != nil {
		return err
	}
	return nil
}

// Execute performs the user retrieval operation.
// It builds the filters from the request and passes them to the user contract to fetch the data.
func (u *ListUserUseCase) Execute() (*models.PaginatedResponse[*models.UserData], *models.SystemError) {
	query := u.request.Build()
	paginatedData, err := u.userContract.GetByFilter(query)
	if err != nil {
		return nil, err
	}

	var result []*models.UserData
	for i := range paginatedData.Rows {
		result = append(result, &models.UserData{
			Id:       paginatedData.Rows[i].ID,
			Username: paginatedData.Rows[i].Username,
			Name:     paginatedData.Rows[i].Name,
			LastName: paginatedData.Rows[i].LastName,
			Email:    paginatedData.Rows[i].Email,
			Type:     paginatedData.Rows[i].Type,
			Picture:  paginatedData.Rows[i].Picture,
			Role:     paginatedData.Rows[i].Role,
			Active:   paginatedData.Rows[i].Active,
		})
	}

	return &models.PaginatedResponse[*models.UserData]{
		TotalRows:  paginatedData.TotalRows,
		TotalPages: paginatedData.TotalPages,
		Rows:       result,
	}, nil
}
