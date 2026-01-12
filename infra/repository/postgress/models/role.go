package models

import (
	"time"

	"hrms.local/core/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RoleGorm struct {
	ID          uuid.UUID        `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Name        string           `gorm:"type:varchar(255);unique;not null"`
	Description string           `gorm:"type:text"`
	Permissions []PermissionGorm `gorm:"foreignKey:RoleID"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func (RoleGorm) TableName() string {
	return "roles"
}

func (r RoleGorm) ToModel() models.Role {
	permissions := make([]models.Permission, len(r.Permissions))
	for i, p := range r.Permissions {
		permissions[i] = p.ToModel()
	}

	return models.Role{
		ID:          r.ID.String(),
		Name:        r.Name,
		Description: r.Description,
		Permissions: permissions,
	}
}

func RoleToEntity(r models.Role) RoleGorm {
	id, _ := uuid.Parse(r.ID)

	permissions := make([]PermissionGorm, len(r.Permissions))
	for i, p := range r.Permissions {
		permissions[i] = PermissionToEntity(p)
	}

	return RoleGorm{
		ID:          id,
		Name:        r.Name,
		Description: r.Description,
		Permissions: permissions,
	}
}
