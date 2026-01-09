package repo

import (
	"hrms.local/core/contracts"
	"hrms.local/core/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DepartmentGorm struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Name      string    `gorm:"type:varchar(255)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (DepartmentGorm) TableName() string {
	return "departments"
}

func (d DepartmentGorm) ToModel() models.Department {
	return models.Department{
		ID:   fromGUIDToString(d.ID),
		Name: d.Name,
	}
}

func ToEntity(d models.Department) DepartmentGorm {
	return DepartmentGorm{
		// Id:   uuid.FromStringOrNil(d.ID),
		Name: d.Name,
	}
}

func fromGUIDToString(id uuid.UUID) string {
	return id.String()
}

type DepartmentRepository struct {
	GenericCrud[models.Department, DepartmentGorm]
}

func NewDepartmentRepository(db *gorm.DB) contracts.DepartmentContract {
	return &DepartmentRepository{
		GenericCrud: NewGenericCrud(db, ToEntity, (DepartmentGorm).ToModel),
	}
}

func (d *DepartmentRepository) SomeMethod() models.SystemError {
	return models.SystemError{}
}
