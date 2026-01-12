package repo

import (
	"hrms.local/core/contracts"
	"hrms.local/core/models"
	gormModels "hrms.local/repository/postgress/models"

	"gorm.io/gorm"
)

type PermissionRepository struct {
	db *gorm.DB
}

func NewPermissionRepository(db *gorm.DB) contracts.PermissionContract {
	return &PermissionRepository{
		db: db,
	}
}

func (p *PermissionRepository) GetAll() ([]models.Permission, *models.SystemError) {
	var permissions []gormModels.PermissionGorm
	if err := p.db.Find(&permissions).Error; err != nil {
		return nil, models.NewSystemError(models.SystemErrorCodeValidation, models.SystemErrorTypeValidation, models.SystemErrorLevelError, "Failed to get permissions", struct{}{})
	}

	result := make([]models.Permission, len(permissions))
	for i, perm := range permissions {
		result[i] = perm.ToModel()
	}

	return result, nil
}
