package repo

import (
	"time"

	"gorm.io/gorm"
	"hrms.local/core/models"
)

type PositionGorm struct {
	ID           string          `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Title        string          `gorm:"type:varchar(255)"`
	Code         string          `gorm:"type:varchar(255)"`
	Description  string          `gorm:"type:text"`
	SalaryMin    float32         `gorm:"type:decimal(10,2)"`
	SalaryMax    float32         `gorm:"type:decimal(10,2)"`
	Currency     string          `gorm:"type:varchar(255)"`
	WorkType     string          `gorm:"type:varchar(255)"`
	DepartmentID string          `gorm:"type:varchar(255)"`
	CreatedAt    time.Time       `gorm:"autoCreateTime"`
	UpdatedAt    time.Time       `gorm:"autoUpdateTime"`
	CreatedByID  uint            `gorm:"type:integer"`
	UpdatedByID  uint            `gorm:"type:integer"`
	Department   *DepartmentGorm `gorm:"foreignKey:DepartmentID"`
	User         *UserGorm       `gorm:"foreignKey:UpdatedByID"`
	Status       string          `gorm:"type:varchar(255)"`
	DeletedAt    gorm.DeletedAt  `gorm:"index"`
}

func (t PositionGorm) TableName() string {
	return "positions"
}

func (t PositionGorm) ToModel() models.Position {
	var dept *models.Department
	if t.Department != nil {
		d := t.Department.ToModel()
		dept = &d
	}

	var user *models.User
	if t.User != nil {
		u := ToEntityUser(*t.User)
		user = &u
	}

	return models.Position{
		ID:           t.ID,
		Title:        t.Title,
		Code:         t.Code,
		Description:  t.Description,
		SalaryMin:    t.SalaryMin,
		SalaryMax:    t.SalaryMax,
		Currency:     t.Currency,
		WorkType:     models.WotkType(t.WorkType),
		DepartmentID: t.DepartmentID,
		CreatedAt:    t.CreatedAt,
		UpdatedAt:    t.UpdatedAt,
		CreatedByID:  t.CreatedByID,
		UpdatedByID:  t.UpdatedByID,
		Department:   dept,
		User:         user,
		Status:       models.PositionStatus(t.Status),
	}
}

type PositionRepository struct {
	GenericCrud[models.Position, PositionGorm]
	db *gorm.DB
}

func PositionGormToModel(t PositionGorm) models.Position {
	return t.ToModel()
}

func PositionGormToEntity(t models.Position) PositionGorm {
	var dept *DepartmentGorm
	if t.Department != nil {
		d := ToEntity(*t.Department)
		dept = &d
	}

	var user *UserGorm
	if t.User != nil {
		u := ToModel(*t.User)
		user = &u
	}

	return PositionGorm{
		ID:           t.ID,
		Title:        t.Title,
		Code:         t.Code,
		Description:  t.Description,
		SalaryMin:    t.SalaryMin,
		SalaryMax:    t.SalaryMax,
		Currency:     t.Currency,
		WorkType:     string(t.WorkType),
		DepartmentID: t.DepartmentID,
		CreatedAt:    t.CreatedAt,
		UpdatedAt:    t.UpdatedAt,
		CreatedByID:  t.CreatedByID,
		UpdatedByID:  t.UpdatedByID,
		Department:   dept,
		User:         user,
		Status:       string(t.Status),
	}
}

func NewPositionRepository(db *gorm.DB) PositionRepository {
	return PositionRepository{
		GenericCrud: NewGenericCrud(db, PositionGormToEntity, PositionGormToModel),
		db:          db,
	}
}
