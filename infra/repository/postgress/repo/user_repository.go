package repo

import (
	"hrms/core/contracts"
	"hrms/core/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserGorm struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Username  string    `gorm:"type:varchar(255)"`
	Password  string    `gorm:"type:varchar(255)"`
	Email     string    `gorm:"type:varchar(255)"`
	Type      string    `gorm:"type:varchar(255)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Active    bool           `gorm:"type:boolean"`
}

func (UserGorm) TableName() string {
	return "users"
}

func ToModel(entity *models.User) *UserGorm {
	return &UserGorm{
		// Id:       uuid.FromStringOrNil(entity.ID),
		Username: entity.Username,
		Password: entity.Password,
		Email:    entity.Email,
		Type:     string(entity.Type),
		Active:   entity.Active,
	}
}

func ToEntityUser(gorm *UserGorm) *models.User {
	return &models.User{
		ID:       fromGUIDToString(gorm.ID),
		Username: gorm.Username,
		Password: gorm.Password,
		Email:    gorm.Email,
		Type:     models.UserType(gorm.Type),
		Active:   gorm.Active,
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
