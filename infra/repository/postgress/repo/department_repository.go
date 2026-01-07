package repo

import (
	"hrms/core/contracts"
	"hrms/core/models"

	"gorm.io/gorm"
)

type DepartmentGorm struct {
	gorm.Model
	Id        string         `gorm:"type:varchar(255);primarykey"`
	Name      string         `gorm:"type:varchar(255)"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (d *DepartmentGorm) ToModel() *models.Department {
	return &models.Department{
		ID:   d.Id,
		Name: d.Name,
	}
}

func ToEntity(d *models.Department) *DepartmentGorm {
	return &DepartmentGorm{
		Id:   d.ID,
		Name: d.Name,
	}
}

type DepartmentRepository struct {
	GenericCrud[models.Department]
}

func NewDepartmentRepository(db *gorm.DB) contracts.DepartmentContract {
	return &DepartmentRepository{
		GenericCrud: NewGenericCrud[models.Department](db),
	}
}

func (d *DepartmentRepository) SomeMethod() models.SystemError {
	return models.SystemError{}
}
