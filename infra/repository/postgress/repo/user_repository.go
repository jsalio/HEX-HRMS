package repo

import (
	"hrms/core/contracts"
	"hrms/core/models"

	"gorm.io/gorm"
)

type UserGorm struct {
	gorm.Model
	Id        string         `gorm:"type:varchar(255);primarykey"`
	Username  string         `gorm:"type:varchar(255)"`
	Password  string         `gorm:"type:varchar(255)"`
	Email     string         `gorm:"type:varchar(255)"`
	Type      string         `gorm:"type:varchar(255)"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func ToModel(entity *models.User) *UserGorm {
	return &UserGorm{
		Id:       entity.ID,
		Username: entity.Username,
		Password: entity.Password,
		Email:    entity.Email,
		Type:     string(entity.Type),
	}
}

func ToEntityUser(gorm *UserGorm) *models.User {
	return &models.User{
		ID:       gorm.Id,
		Username: gorm.Username,
		Password: gorm.Password,
		Email:    gorm.Email,
		Type:     models.UserType(gorm.Type),
	}
}

type UserRepository struct {
	GenericCrud[models.User]
}

func NewUserRepository(db *gorm.DB) contracts.UserContract {
	return &UserRepository{
		GenericCrud: NewGenericCrud[models.User](db),
	}
}
