package user

import (
	"hrms/core/contracts"
	"hrms/core/models"
)

// ModifyUserUseCase handles the modification of an existing user
// It orchestrates the validation of the request and the updating of the user via the user contract
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
//	request := &contracts.GenericRequest[models.ModifyUser]{Data: models.ModifyUser{
//		ID:       "123",
//		Username: "jdoe",
//		Password: "password",
//		Email:    "jdoe@example.com",
//		Type:     models.UserTypeAdmin,
//	}}
//
//	// 3. Instantiate the UseCase
//	useCase := user.NewModifyUserUseCase(userRepo, request)
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
type ModifyUserUseCase struct {
	userContract contracts.UserContract
	request      contracts.IGenericRequest[models.ModifyUser]
}

// NewModifyUserUseCase creates a new instance of ModifyUserUseCase
// It injects the user contract (dependency inversion) and the request data
func NewModifyUserUseCase(userContract contracts.UserContract, request contracts.IGenericRequest[models.ModifyUser]) *ModifyUserUseCase {
	return &ModifyUserUseCase{
		userContract: userContract,
		request:      request,
	}
}

// Validate ensures that the request data is valid
// It checks if the request is empty and validates the user data using reflection
func (u *ModifyUserUseCase) Validate() *models.SystemError {
	request := u.request.Build()
	if err := request.Validate(); err != nil {
		return err
	}
	return nil
}

// Execute performs the user modification operation
// It builds the user data from the request and passes it to the user contract to update the user
func (u *ModifyUserUseCase) Execute() (*string, *models.SystemError) {
	request := u.request.Build()
	user, err := u.userContract.Update(request.ID, *request.ToUser())
	if err != nil {
		return nil, err
	}
	result := models.NewDynamicResult(user)
	data := models.MustConvert[string](result)
	return &data, nil
}
