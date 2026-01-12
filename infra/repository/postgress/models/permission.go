package models

import (
	"hrms.local/core/models"

	"github.com/google/uuid"
)

type PermissionGorm struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Name        string    `gorm:"type:varchar(255);unique;not null"`
	Description string    `gorm:"type:text"`
	RoleID      uuid.UUID `gorm:"type:uuid;not null"`
	Role        RoleGorm  `gorm:"foreignKey:RoleID"`
}

func (PermissionGorm) TableName() string {
	return "permissions"
}

func (p PermissionGorm) ToModel() models.Permission {
	return models.Permission{
		ID:          p.ID.String(),
		Name:        p.Name,
		Description: p.Description,
		RoleId:      p.RoleID.String(),
	}
}

func PermissionToEntity(p models.Permission) PermissionGorm {
	id, _ := uuid.Parse(p.ID)
	roleID, _ := uuid.Parse(p.RoleId)
	return PermissionGorm{
		ID:          id,
		Name:        p.Name,
		Description: p.Description,
		RoleID:      roleID,
	}
}
