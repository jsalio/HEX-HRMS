package contracts

import (
	"hrms.local/core/models"
)

// define basic user operations
// example :
//
//	userContract := NewUserContract()
//	userContract.Create(models.User{ID: "1", Username: "HR", Password: "HR", Email: "HR", Type: "HR"})
//	userContract.Update("1", models.User{ID: "1", Username: "HR", Password: "HR", Email: "HR", Type: "HR"})
//	userContract.Delete("1")
//	userContract.GetAll()
//	userContract.GetByFilter(models.Filter{Key: "name", Value: "HR"})
type UserContract interface {
	ReadOperation[models.User]
	WriteOperation[models.User]
}
