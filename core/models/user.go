package models

type UserType string

const (
	UserTypeAdmin  UserType = "admin"
	UserTypeNormal UserType = "normal"
)

type User struct {
	ID       string
	Username string
	Name     string
	LastName string
	Password string
	Email    string
	Type     UserType
	Active   bool
	Picture  string
	Role     string
}

type CreateUser struct {
	Username string
	Password string
	Email    string
	Role     string
	Name     string
	LastName string
	Type     UserType
	Picture  string
}

type ModifyUser struct {
	ID       string
	Username string
	Name     string
	LastName string
	Password string
	Email    string
	Type     UserType
}

type UserData struct {
	Id       string   `json:"id"`
	Username string   `json:"username"`
	Name     string   `json:"name"`
	LastName string   `json:"lastName"`
	Email    string   `json:"email"`
	Type     UserType `json:"type"`
	Picture  string   `json:"picture"`
	Role     string   `json:"role"`
	Active   bool     `json:"active"`
}

type LoginUser struct {
	Username string
	Password string
}

func (u *User) ToUserData() *UserData {
	return &UserData{
		Id:       u.ID,
		Username: u.Username,
		Name:     u.Name,
		LastName: u.LastName,
		Email:    u.Email,
		Type:     u.Type,
		Picture:  u.Picture,
		Role:     u.Role,
	}
}

func (cu *CreateUser) ToUser() *User {
	return &User{
		Username: cu.Username,
		Password: cu.Password,
		Email:    cu.Email,
		Type:     cu.Type,
		Name:     cu.Name,
		LastName: cu.LastName,
		Picture:  cu.Picture,
		Role:     cu.Role,
		Active:   true,
	}
}

func (mu *ModifyUser) ToUser() *User {
	return &User{
		ID:       mu.ID,
		Username: mu.Username,
		Name:     mu.Name,
		LastName: mu.LastName,
		Password: mu.Password,
		Email:    mu.Email,
		Type:     mu.Type,
	}
}

func (cu *CreateUser) Validate() *SystemError {
	if cu.Username == "" {
		return NewSystemError(SystemErrorCodeValidation, SystemErrorTypeValidation, SystemErrorLevelError, "username is required", struct{}{})
	}
	if cu.Password == "" {
		return NewSystemError(SystemErrorCodeValidation, SystemErrorTypeValidation, SystemErrorLevelError, "password is required", struct{}{})
	}
	if cu.Email == "" {
		return NewSystemError(SystemErrorCodeValidation, SystemErrorTypeValidation, SystemErrorLevelError, "email is required", struct{}{})
	}
	if cu.Type == "" {
		return NewSystemError(SystemErrorCodeValidation, SystemErrorTypeValidation, SystemErrorLevelError, "type is required", struct{}{})
	}
	return nil
}

func (mu *ModifyUser) Validate() *SystemError {
	if mu.ID == "" {
		return NewSystemError(SystemErrorCodeValidation, SystemErrorTypeValidation, SystemErrorLevelError, "id is required", struct{}{})
	}
	if mu.Username == "" {
		return NewSystemError(SystemErrorCodeValidation, SystemErrorTypeValidation, SystemErrorLevelError, "username is required", struct{}{})
	}
	if mu.Password == "" {
		return NewSystemError(SystemErrorCodeValidation, SystemErrorTypeValidation, SystemErrorLevelError, "password is required", struct{}{})
	}
	if mu.Email == "" {
		return NewSystemError(SystemErrorCodeValidation, SystemErrorTypeValidation, SystemErrorLevelError, "email is required", struct{}{})
	}
	if mu.Type == "" {
		return NewSystemError(SystemErrorCodeValidation, SystemErrorTypeValidation, SystemErrorLevelError, "type is required", struct{}{})
	}
	if mu.Name == "" {
		return NewSystemError(SystemErrorCodeValidation, SystemErrorTypeValidation, SystemErrorLevelError, "name is required", struct{}{})
	}
	if mu.LastName == "" {
		return NewSystemError(SystemErrorCodeValidation, SystemErrorTypeValidation, SystemErrorLevelError, "last name is required", struct{}{})
	}
	return nil
}
