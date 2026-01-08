package user

import (
	"hrms/core/contracts"
	"hrms/core/models"
)

type LoginUserUseCase struct {
	userContract contracts.UserContract
	request      contracts.IGenericRequest[models.LoginUser]
}

func NewLoginUserUseCase(userContract contracts.UserContract, request contracts.IGenericRequest[models.LoginUser]) *LoginUserUseCase {
	return &LoginUserUseCase{
		userContract: userContract,
		request:      request,
	}
}

func (u *LoginUserUseCase) Validate() *models.SystemError {
	request := u.request.Build()
	if request.Username == "" || request.Password == "" {
		return models.NewSystemError(models.SystemErrorCodeValidation, models.SystemErrorTypeValidation, models.SystemErrorLevelError, "request is empty", struct{}{})
	}
	user, err := u.userContract.GetByFilter(models.Filter{
		Key:   "username",
		Value: request.Username,
	})
	if err != nil {
		return err
	}
	if user == nil {
		return models.NewSystemError(models.SystemErrorCodeInternal, models.SystemErrorTypeValidation, models.SystemErrorLevelError, "El usuario no existe", struct{}{})
	}

	if user[0].Password != request.Password {
		return models.NewSystemError(models.SystemErrorCodeInternal, models.SystemErrorTypeValidation, models.SystemErrorLevelError, "Contraseña incorrecta", struct{}{})
	}

	return nil
}

func (u *LoginUserUseCase) Execute() (*models.UserData, *models.SystemError) {
	request := u.request.Build()
	user, err := u.userContract.GetByFilter(models.Filter{
		Key:   "username",
		Value: request.Username,
	})
	if err != nil {
		return nil, err
	}
	return user[0].ToUserData(), nil
}
