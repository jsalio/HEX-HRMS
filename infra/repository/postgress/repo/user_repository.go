package repo

import (
	"time"

	"hrms.local/core/contracts"
	"hrms.local/core/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserGorm struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Username  string    `gorm:"type:varchar(255)"`
	Password  string    `gorm:"type:varchar(255)"`
	Email     string    `gorm:"type:varchar(255)"`
	Name      string    `gorm:"type:varchar(255)"`
	LastName  string    `gorm:"type:varchar(255)"`
	Type      string    `gorm:"type:varchar(255)"`
	Picture   string    `gorm:"type:varchar(255)"`
	Role      string    `gorm:"type:varchar(255)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Active    bool           `gorm:"type:boolean"`
}

func (UserGorm) TableName() string {
	return "users"
}

func ToModel(entity models.User) UserGorm {
	id := uuid.Nil
	if entity.ID != "" {
		id, _ = uuid.Parse(entity.ID)
	}
	return UserGorm{
		ID:       id,
		Username: entity.Username,
		Password: entity.Password,
		Email:    entity.Email,
		Name:     entity.Name,
		LastName: entity.LastName,
		Type:     string(entity.Type),
		Picture:  entity.Picture,
		Role:     string(entity.Role),
		Active:   entity.Active,
	}
}

func ToEntityUser(gorm UserGorm) models.User {
	return models.User{
		ID:       fromGUIDToString(gorm.ID),
		Username: gorm.Username,
		Password: gorm.Password,
		Email:    gorm.Email,
		Name:     gorm.Name,
		LastName: gorm.LastName,
		Type:     models.UserType(gorm.Type),
		Picture:  gorm.Picture,
		Role:     gorm.Role,
		Active:   gorm.Active,
	}
}

type UserRepository struct {
	GenericCrud[models.User, UserGorm]
}

func (r *UserRepository) Update(id string, item models.User) (models.User, *models.SystemError) {
	if item.Password == "" {
		existing, err := r.GetOnce("id", id)
		if err == nil && existing != nil {
			item.Password = existing.Password
		}
	}
	return r.GenericCrud.Update(id, item)
}

func NewUserRepository(db *gorm.DB) contracts.UserContract {
	return &UserRepository{
		GenericCrud: NewGenericCrud(db, ToModel, ToEntityUser),
	}
}
