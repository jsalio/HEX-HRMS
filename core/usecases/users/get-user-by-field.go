package user

import (
	"hrms.local/core/contracts"
	"hrms.local/core/models"
)

type GetUserByFieldUseCase struct {
	userContract contracts.UserContract
	request      contracts.IGenericRequest[models.Filter]
}

func NewGetUserByFieldUseCase(userContract contracts.UserContract, request contracts.IGenericRequest[models.Filter]) *GetUserByFieldUseCase {
	return &GetUserByFieldUseCase{
		userContract: userContract,
		request:      request,
	}
}

func (u *GetUserByFieldUseCase) Validate() *models.SystemError {
	request := u.request.Build()
	if request.Key == "" || request.Value == "" {
		return models.NewSystemError(models.SystemErrorCodeValidation, models.SystemErrorTypeValidation, models.SystemErrorLevelError, "Key and Value are required", nil)
	}
	if request.Key != "username" && request.Key != "email" && request.Key != "id" {
		return models.NewSystemError(models.SystemErrorCodeValidation, models.SystemErrorTypeValidation, models.SystemErrorLevelError, "Key must be username, email or id", nil)
	}
	exists, err := u.userContract.Exists(request.Key, request.Value)
	if err != nil {
		return models.NewSystemError(models.SystemErrorCodeInternal, models.SystemErrorTypeInternal, models.SystemErrorLevelError, "Error checking if user exists", nil)
	}
	if !exists {
		return models.NewSystemError(models.SystemErrorCodeNone, models.SystemErrorType(models.SystemErrorLevelError), models.SystemErrorLevelError, "User not found", nil)
	}

	return nil
}

func (u *GetUserByFieldUseCase) Execute() (*models.UserData, *models.SystemError) {
	request := u.request.Build()
	user, err := u.userContract.GetOnce(request.Key, request.Value)
	if err != nil {
		return nil, models.NewSystemError(models.SystemErrorCodeInternal, models.SystemErrorTypeInternal, models.SystemErrorLevelError, "Error getting user", nil)
	}
	result := &models.UserData{
		Id:       user.ID,
		Username: user.Username,
		Name:     user.Name,
		LastName: user.LastName,
		Email:    user.Email,
		Type:     user.Type,
		Picture:  user.Picture,
		Role:     user.Role,
		Active:   user.Active,
	}
	return result, nil
}
