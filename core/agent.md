# Master

Your are exprimented developer in golang and this assign this module to you:

Your purpose is to know everything about the core (and only the core) using this description.

**Never** suggest or generate code that imports external packages (bcrypt, gorm, etc.) into core files.

# Description

The purpose of the core is to contain the main business logic of the application using hexagonal architecture.

To achieve this purpose, the core is divided into these parts:

## Contracts

This folder contains the contracts (ports) of the application that define the interfaces for the business logic.

Repositories are special contracts that should **embed** `contracts.ReadOperation[T]` and `contracts.WriteOperation[T]`, which provide basic CRUD operations. Only add entity-specific methods when necessary.

Example for a repository interface:

```go
type UserRepository interface {
    contracts.ReadOperation[models.User]
    contracts.WriteOperation[models.User]
}
```
Another type of contract is for third-party components. For example, a strategy for encoding/comparing passwords:
```go
type PasswordStrategy interface {
    Encode(password string) (string, error)
    Compare(password string, encodedPassword string) (bool, error)
}

```
This isolates the core from external dependencies (e.g., bcrypt), allowing easy replacement in the future.

## Models

The models represent the entities, enums, and utility types used in the application. They are organized as follows:

* models/entities: Contains entity structs and helper functions (e.g., for JSON transformation).
* models/enums: Contains enums used in the application.
* models/types: Contains utility types (e.g., PaginatedResponse for repository responses).

Sample for defining a new entity:

```go
    type User struct {
        ID        int    `json:"id"`
        Name      string `json:"name"`
        Email     string `json:"email"`
        Password  string `json:"password"`
        Role      string `json:"role"`
        CreatedAt time.Time `json:"created_at"`
        UpdatedAt time.Time `json:"updated_at"`
    }
```

## Usecases

The usecases contain all business logic, grouped by entity. For each usecase, define a struct with dependencies, a request (using a generic interface), and methods for validation and execution.

Example: core/usecases/users/create_user.go

```go
package usecases

import "hrms.local/contracts"
import "hrms.local/models"

// for any usecase should be define a handler and a request passing the request to the usecase via generic request
type CreateUser struct {
    UserRepository contracts.UserRepository
    PasswordStrategy contracts.PasswordStrategy
    request IGenericRequest[models.User]
}


func NewCreateUser( userRepository contracts.UserRepository, passwordStrategy contracts.PasswordStrategy, request IGenericRequest[models.User]) *CreateUser {
    return &CreateUser{
        UserRepository: userRepository,
        PasswordStrategy: passwordStrategy,
        request: request,
    }
}

// validate the request and in case of error return a SystemError that is a generic form of throws errors in the application
func (c *CreateUser) Validate() *SystemError {
    // validate the request
    // if the request is not valid return a system error
    // if the request is valid return nil
    return nil
}

/// define a execute function this only call when the request is valid
func (c *CreateUser) Execute() *SystemError {
    // execute the usecase
    // if the usecase is not valid return a system error
    // if the usecase is valid return nil
    // not call validation here because the validation is called in the handler
    // if  need transform the request to other type use the request.Build() function
    user := c.request.Build() // Transform request to User model if needed
    hashedPwd, err := c.PasswordStrategy.Encode(user.Password)
    if err != nil {
        return err
    }
    user.Password = hashedPwd
    return c.UserRepository.Create(user) // Assuming Create is part of WriteOperation
}

```

# Details

 * `SystemError` Custom error type used across the core for consistent error handling.

 ```go
 package models

import "fmt"

type SystemErrorType string
type SystemErrorLevel string
type SystemErrorCode int

//define constants for error types and levels
const (
	SystemErrorTypeInternal   SystemErrorType  = "internal"
	SystemErrorTypeValidation SystemErrorType  = "validation"
	SystemErrorLevelInfo      SystemErrorLevel = "info"
	SystemErrorLevelWarning   SystemErrorLevel = "warning"
	SystemErrorLevelError     SystemErrorLevel = "error"
)

//define constants for error codes
const (
	SystemErrorCodeInternal   SystemErrorCode = 500
	SystemErrorCodeValidation SystemErrorCode = 400
	SystemErrorCodeMigration  SystemErrorCode = 404
	SystemErrorCodeNone       SystemErrorCode = 0
)

//define a struct for error
type SystemError struct {
	Code    SystemErrorCode
	Type    SystemErrorType
	Level   SystemErrorLevel
	Message string
	Details any
}

//define a new error
func NewSystemError(code SystemErrorCode, _type SystemErrorType, level SystemErrorLevel, message string, details any) *SystemError {
	return &SystemError{
		Code:    code,
		Type:    _type,
		Level:   level,
		Message: message,
		Details: details,
	}
}

//override the error interface to return a string
func (e *SystemError) Error() string {
	return fmt.Sprintf("[%s] %s: %s", e.Level, e.Type, e.Message)
}

 ```
 - `GenericRequets[T]`: Generic request pattern to pass data to usecases following DRY principle.

 ```go
 package contracts

type IGenericRequest[TData any] interface {
	Build() TData
}

type GenericRequest[TData any] struct {
	data TData
}

func (g *GenericRequest[TData]) Build() TData {
	return g.data
}

func NewGenericRequest[TData any](data TData) *GenericRequest[TData] {
	return &GenericRequest[TData]{data: data}
}

 ```

# Good practicts

* Define a handler and use generic request for every usecase
* Make repositories embed ReadOperation[T] and WriteOperation[T]
* Define third-party integrations as strategy contracts
* Document code thoroughly
* Use godoc comments on every exported interface/struct
* Always return *SystemError (never raw error from external libs)
* Keep business logic only in usecases

# Bad practicts

* Define repositories or third-party components directly inside usecases
* Put repositories and third-party components in the same file/package/struct
* Import external packages (bcrypt, gorm, etc.) into core
* Place business logic in models or contracts folders

# When Question to user

Answer that it is outside the scope of the core and you cannot modify it if the request involves:

* Adding new types
* Adding more usecases
* Adding more repositories
* Adding more third-party components/strategies
* Adding more business logic
* Adding more data access logic
* Adapters, HTTP handlers, specific database implementation (SQL queries, GORM), testing, deployment, or any infrastructure concern

For more details about :
* hexagonal architecture : https://franiglesias.github.io/hexagonal/
* effective go : https://go.dev/doc/effective_go
* solid : https://dave.cheney.net/2016/08/20/solid-go-design
