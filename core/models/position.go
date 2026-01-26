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

// Position model represents a job position request
type Position struct {
	ID             string
	Title          string
	Code           string
	Description    string
	RequiredSkills string // Skills required for this position (text)
	SalaryMin      float32
	SalaryMax      float32
	MaxEmployees   int // Maximum number of employees to hire
	Currency       string
	WorkType       WotkType
	DepartmentID   string
	Department     *Department
	CreatedAt      time.Time
	UpdatedAt      time.Time
	CreatedByID    uint
	UpdatedByID    uint
	User           *User
	Status         PositionStatus
}

// CreatePosition represents data needed to create a new position
type CreatePosition struct {
	Title          string
	Code           string
	Description    string
	RequiredSkills string
	SalaryMin      float32
	SalaryMax      float32
	MaxEmployees   int
	Currency       string
	WorkType       WotkType
	DepartmentID   string
	CreatedByID    uint
}

// ModifyPosition represents data needed to update an existing position
type ModifyPosition struct {
	ID             string
	Title          string
	Code           string
	Description    string
	RequiredSkills string
	SalaryMin      float32
	SalaryMax      float32
	MaxEmployees   int
	Currency       string
	WorkType       WotkType
	DepartmentID   string
	UpdatedByID    uint
	Status         PositionStatus
}

// PositionItem represents a position for list views
type PositionItem struct {
	ID             string
	Title          string
	Code           string
	Description    string
	RequiredSkills string
	SalaryMin      float32
	SalaryMax      float32
	MaxEmployees   int
	Currency       string
	WorkType       WotkType
	DepartmentID   string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	CreatedByID    uint
	UpdatedByID    uint
	Status         PositionStatus
}

func (cp *CreatePosition) ToPosition() *Position {
	return &Position{
		Title:          cp.Title,
		Code:           cp.Code,
		Description:    cp.Description,
		RequiredSkills: cp.RequiredSkills,
		SalaryMin:      cp.SalaryMin,
		SalaryMax:      cp.SalaryMax,
		MaxEmployees:   cp.MaxEmployees,
		Currency:       cp.Currency,
		WorkType:       cp.WorkType,
		DepartmentID:   cp.DepartmentID,
		CreatedByID:    cp.CreatedByID,
		Status:         Active,
	}
}

func (mp *ModifyPosition) ToPosition() *Position {
	return &Position{
		ID:             mp.ID,
		Title:          mp.Title,
		Code:           mp.Code,
		Description:    mp.Description,
		RequiredSkills: mp.RequiredSkills,
		SalaryMin:      mp.SalaryMin,
		SalaryMax:      mp.SalaryMax,
		MaxEmployees:   mp.MaxEmployees,
		Currency:       mp.Currency,
		WorkType:       mp.WorkType,
		DepartmentID:   mp.DepartmentID,
		UpdatedByID:    mp.UpdatedByID,
		Status:         mp.Status,
	}
}

func (cp *Position) ToPositionItem() *PositionItem {
	return &PositionItem{
		ID:             cp.ID,
		Title:          cp.Title,
		Code:           cp.Code,
		Description:    cp.Description,
		RequiredSkills: cp.RequiredSkills,
		SalaryMin:      cp.SalaryMin,
		SalaryMax:      cp.SalaryMax,
		MaxEmployees:   cp.MaxEmployees,
		Currency:       cp.Currency,
		WorkType:       cp.WorkType,
		DepartmentID:   cp.DepartmentID,
		CreatedAt:      cp.CreatedAt,
		UpdatedAt:      cp.UpdatedAt,
		CreatedByID:    cp.CreatedByID,
		UpdatedByID:    cp.UpdatedByID,
		Status:         cp.Status,
	}
}
