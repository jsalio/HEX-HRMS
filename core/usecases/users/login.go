package user

import (
	"hrms.local/core/contracts"
	"hrms.local/core/models"
)

type LoginUserUseCase struct {
	userContract         contracts.UserContract
	request              contracts.IGenericRequest[models.LoginUser]
	cryptographyContract contracts.CryptographyContract
}

func NewLoginUserUseCase(userContract contracts.UserContract, request contracts.IGenericRequest[models.LoginUser], cryptographyContract contracts.CryptographyContract) *LoginUserUseCase {
	return &LoginUserUseCase{
		userContract:         userContract,
		request:              request,
		cryptographyContract: cryptographyContract,
	}
}

func (u *LoginUserUseCase) Validate() *models.SystemError {
	request := u.request.Build()
	if request.Username == "" || request.Password == "" {
		return models.NewSystemError(models.SystemErrorCodeValidation, models.SystemErrorTypeValidation, models.SystemErrorLevelError, "request is empty", struct{}{})
	}
	user, err := u.userContract.GetByFilter(models.SearchQuery{
		Filters: models.Filters{
			{
				Key:   "Username",
				Value: request.Username,
			},
		},
	})
	if err != nil {
		return err
	}
	if len(user) == 0 {
		return models.NewSystemError(models.SystemErrorCodeInternal, models.SystemErrorTypeValidation, models.SystemErrorLevelError, "El usuario no existe", struct{}{})
	}

	// Compare the plain text password with the hashed password
	isValid, err := u.cryptographyContract.ComparePassword(request.Password, user[0].Password)
	if err != nil {
		return err
	}
	if !isValid {
		return models.NewSystemError(models.SystemErrorCodeInternal, models.SystemErrorTypeValidation, models.SystemErrorLevelError, "Contraseña incorrecta", struct{}{})
	}

	return nil
}

func (u *LoginUserUseCase) Execute() (*models.UserData, *models.SystemError) {
	request := u.request.Build()
	user, err := u.userContract.GetByFilter(models.SearchQuery{
		Filters: models.Filters{
			{
				Key:   "Username",
				Value: request.Username,
			},
		},
	})
	if err != nil {
		return nil, err
	}
	return user[0].ToUserData(), nil
}
