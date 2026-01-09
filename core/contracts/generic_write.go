package contracts

import "hrms.local/core/models"

// define basic write operations
type WriteOperation[T any] interface {
	// Create basic resource in repository
	// example :
	// 	 data,err:=Create(item models.Department)
	// 	 if err != nil {
	// 		return nil, err
	// 	}
	// 	return data, nil
	Create(item T) (T, *models.SystemError)
	// Update basic resource in repositor	y
	// example :
	// 	 data,err:=Update(id string, item models.Department)
	// 	 if err != nil {
	// 		return nil, err
	// 	}
	// 	return data, nil
	Update(id string, item T) (interface{}, *models.SystemError)
	// Delete basic resource in repository
	// example :
	// 		data,err:=Delete(id string)
	// 		if err != nil {
	// 			return nil, err
	// 		}
	// 		return data, nil
	Delete(id string) (interface{}, error)
}
