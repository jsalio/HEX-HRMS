package user

import (
	"hrms/core/contracts"
	"hrms/core/models"
)

// CreateUserUseCase handles the creation of a new user.
// It orchestrates the validation of the user data and the creation of the user via the user contract.
//
// Example Usage:
//
//	// 1. Create the repository/contract implementation
//	var userRepo contracts.UserContract = ... // Your repository implementation
//
//	// 2. Create the request with user data
//	user := models.CreateUser{
//		Username: "jdoe",
//		Password: "password123",
//		Email:    "jdoe@example.com",
//		Type:     models.UserTypeNormal,
//	}
//	// Assuming you have a concrete implementation of IGenericRequest or use the struct directly if generic
//	request := &contracts.GenericRequest[models.CreateUser]{Data: user}
//
//	// 3. Instantiate the UseCase
//	useCase := user.NewCreateUserUseCase(userRepo, request)
//
//	// 4. Validate the request
//	if err := useCase.Validate(); err != nil {
//	    log.Printf("Validation failed: %v", err.Message)
//	    return
//	}
//
//	// 5. Execute the logic
//	createdUser, err := useCase.Execute()
//	if err != nil {
//	    log.Printf("Execution failed: %v", err.Message)
//	    return
//	}
//
//	fmt.Printf("User created: %s\n", createdUser.Username)
type CreateUserUseCase struct {
	userContract contracts.UserContract
	request      contracts.IGenericRequest[models.CreateUser]
}

// NewCreateUserUseCase creates a new instance of CreateUserUseCase.
// It injects the user contract (dependency inversion) and the request data.
func NewCreateUserUseCase(userContract contracts.UserContract, request contracts.IGenericRequest[models.CreateUser]) *CreateUserUseCase {
	return &CreateUserUseCase{
		userContract: userContract,
		request:      request,
	}
}

// Validate ensures that the user data is valid.
// It checks if the request is empty and validates the user data using reflection.
func (u *CreateUserUseCase) Validate() *models.SystemError {
	request := u.request.Build()

	if err := request.Validate(); err != nil {
		return err
	}

	user, err := u.userContract.GetByFilter(models.Filter{
		Key:   "username",
		Value: request.Username,
	})
	if err != nil {
		return err
	}
	if user != nil {
		return models.NewSystemError(models.SystemErrorCodeInternal, models.SystemErrorTypeValidation, models.SystemErrorLevelError, "El usuario ya existe", struct{}{})
	}
	return nil
}

// Execute performs the user creation operation.
// It builds the user data from the request and passes it to the user contract to create the user.
func (u *CreateUserUseCase) Execute() (*models.UserData, *models.SystemError) {
	request := u.request.Build()
	newUser := request.ToUser()
	newUser.Active = true
	user, err := u.userContract.Create(*newUser)
	if err != nil {
		return nil, err
	}
	return user.ToUserData(), nil
}
