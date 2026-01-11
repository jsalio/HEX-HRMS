package contracts

import "hrms.local/core/models"

// define basic read operations
type ReadOperation[T any] interface {
	// Get resource by filter in repository
	// example :
	// 		data,err:=GetByFilter(models.Filter{Key: "name", Value: "HR"} )
	// 		if err != nil {
	// 			return nil, err
	// 		}
	// 		return data, nil
	GetByFilter(query models.SearchQuery) (*models.PaginatedResponse[T], *models.SystemError)

	// check if resource exists in repository
	// example :
	// 		exists,err:=Exists("username", "HR")
	// 		if err != nil {
	// 			return false, err
	// 		}
	// 		return exists, nil
	Exists(key string, value any) (bool, *models.SystemError)

	// Get resource by field in repository
	// example :
	// 		data,err:=GetOnce("username", "HR")
	// 		if err != nil {
	// 			return nil, err
	// 		}
	// 		return data, nil
	GetOnce(key string, value any) (*T, *models.SystemError)
}
