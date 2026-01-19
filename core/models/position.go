package models

import "time"

type WotkType string

type PositionStatus string

const (
	Active   PositionStatus = "Active"
	Inactive PositionStatus = "Inactive"
	Closed   PositionStatus = "Closed"
)

const (
	Remote WotkType = "Remote"
	Hybrid WotkType = "Hybrid"
	OnSite WotkType = "OnSite"
)

// Position model
type Position struct {
	ID           string
	Title        string
	Code         string
	Description  string
	SalaryMin    float32
	SalaryMax    float32
	Currency     string
	WorkType     WotkType
	DepartmentID string
	Department   *Department
	CreatedAt    time.Time
	UpdatedAt    time.Time
	CreatedByID  uint
	UpdatedByID  uint
	User         *User
	Status       PositionStatus
}

type CreatePosition struct {
	Title        string
	Code         string
	Description  string
	SalaryMin    float32
	SalaryMax    float32
	Currency     string
	WorkType     WotkType
	DepartmentID string
	CreatedByID  uint
}

type ModifyPosition struct {
	ID           string
	Title        string
	Code         string
	Description  string
	SalaryMin    float32
	SalaryMax    float32
	Currency     string
	WorkType     WotkType
	DepartmentID string
	UpdatedByID  uint
}

type PositionItem struct {
	ID           string
	Title        string
	Code         string
	Description  string
	SalaryMin    float32
	SalaryMax    float32
	Currency     string
	WorkType     WotkType
	DepartmentID string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	CreatedByID  uint
	UpdatedByID  uint
}

func (cp *CreatePosition) ToPosition() *Position {
	return &Position{
		Title:        cp.Title,
		Code:         cp.Code,
		Description:  cp.Description,
		SalaryMin:    cp.SalaryMin,
		SalaryMax:    cp.SalaryMax,
		Currency:     cp.Currency,
		WorkType:     cp.WorkType,
		DepartmentID: cp.DepartmentID,
		CreatedByID:  cp.CreatedByID,
	}
}

func (mp *ModifyPosition) ToPosition() *Position {
	return &Position{
		ID:           mp.ID,
		Title:        mp.Title,
		Code:         mp.Code,
		Description:  mp.Description,
		SalaryMin:    mp.SalaryMin,
		SalaryMax:    mp.SalaryMax,
		Currency:     mp.Currency,
		WorkType:     mp.WorkType,
		DepartmentID: mp.DepartmentID,
		UpdatedByID:  mp.UpdatedByID,
	}
}

func (cp *Position) ToPositionItem() *PositionItem {
	return &PositionItem{
		ID:           cp.ID,
		Title:        cp.Title,
		Code:         cp.Code,
		Description:  cp.Description,
		SalaryMin:    cp.SalaryMin,
		SalaryMax:    cp.SalaryMax,
		Currency:     cp.Currency,
		WorkType:     cp.WorkType,
		DepartmentID: cp.DepartmentID,
		CreatedAt:    cp.CreatedAt,
		UpdatedAt:    cp.UpdatedAt,
		CreatedByID:  cp.CreatedByID,
		UpdatedByID:  cp.UpdatedByID,
	}
}
