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
	GetByFilter(query models.SearchQuery) ([]T, *models.SystemError)
}
