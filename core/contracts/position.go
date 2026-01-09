package contracts

import "hrms.local/core/models"

// define basic position operations
// example :
//
//	data,err:=positionContract.GetAll()
//	if err != nil {
//		return nil, err
//	}
//	return data, nil
//
//	data,err:=positionContract.GetByFilter("name", "HR")
//	if err != nil {
//		return nil, err
//	}
//	return data, nil
//
//	data,err:=positionContract.Create(models.Position{
//		ID:   "1",
//		Name: "HR",
//	})
//
//	if err != nil {
//		return nil, err
//	}
//	return data, nil
//	data,err:=positionContract.Update("1", models.Position{
//	ID:   "1",
//	Name: "HR",
//	})
//
//	if err != nil {
//		return nil, err
//	}
//	return data, nil
//
//	data,err:=positionContract.Delete("1")
//	if err != nil {
//		return nil, err
//	}
//	return data, nil
type PositionContract interface {
	// define basic read operations
	ReadOperation[models.Position]
	// define basic write operations
	WriteOperation[models.Position]
}
