package models

// Position model
type Position struct {
	// ID of the position
	ID string
	// Name of the position
	Name string
	// ID of the department
	DepartmentID string
	// Department of the position
	Department *Department
}
