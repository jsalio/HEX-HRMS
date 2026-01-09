package contracts

import "hrms.local/core/models"

// define basic department operations
// example :
//
//	departmentContract := NewDepartmentContract()
//	departmentContract.Create(models.Department{ID: "1", Name: "HR"})
//	departmentContract.Update("1", models.Department{ID: "1", Name: "HR"})
//	departmentContract.Delete("1")
//	departmentContract.GetAll()
//	departmentContract.GetByFilter("name", "HR")
type DepartmentContract interface {
	// define basic read operations
	ReadOperation[models.Department]
	// define basic write operations
	WriteOperation[models.Department]

	SomeMethod() models.SystemError
}
