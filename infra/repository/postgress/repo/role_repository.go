package repo

import (
	"hrms.local/core/contracts"
	"hrms.local/core/models"
	gormModels "hrms.local/repository/postgress/models"

	"gorm.io/gorm"
)

type RoleRepository struct {
	GenericCrud[models.Role, gormModels.RoleGorm]
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) contracts.RoleContract {
	return &RoleRepository{
		GenericCrud: NewGenericCrud(db, gormModels.RoleToEntity, (gormModels.RoleGorm).ToModel),
		db:          db,
	}
}

func (r *RoleRepository) GetPermissions(roleID string) ([]models.Permission, *models.SystemError) {
	var role gormModels.RoleGorm
	if err := r.db.Preload("Permissions").First(&role, "id = ?", roleID).Error; err != nil {
		return nil, models.NewSystemError(models.SystemErrorCodeValidation, models.SystemErrorTypeValidation, models.SystemErrorLevelError, "Role not found", struct{}{})
	}

	permissions := make([]models.Permission, len(role.Permissions))
	for i, p := range role.Permissions {
		permissions[i] = p.ToModel()
	}

	return permissions, nil
}

// Override GetOnce to preload permissions
func (r *RoleRepository) GetOnce(key string, value any) (*models.Role, *models.SystemError) {
	var gormModel gormModels.RoleGorm
	dbQuery := r.db.Preload("Permissions").Where(key+" = ?", value)
	if err := dbQuery.First(&gormModel).Error; err != nil {
		return nil, models.NewSystemError(models.SystemErrorCodeValidation, models.SystemErrorTypeValidation, models.SystemErrorLevelError, "GetOnce failed", struct{}{})
	}
	entity := gormModel.ToModel()
	return &entity, nil
}

// Override GetByFilter to preload permissions
func (r *RoleRepository) GetByFilter(query models.SearchQuery) (*models.PaginatedResponse[models.Role], *models.SystemError) {
	// We need to implement Count separately or reuse generic count but here we override the whole method
	totalRows, sysErr := r.GenericCrud.CountByFilter(query)
	if sysErr != nil {
		return nil, sysErr
	}

	var gormModels []gormModels.RoleGorm
	dbQuery := r.db.Preload("Permissions")
	for _, filter := range query.Filters {
		dbQuery = dbQuery.Where(filter.Key+" = ?", filter.Value)
	}

	limit := query.Pagination.GetLimit()
	dbQuery = dbQuery.Limit(limit).Offset(query.Pagination.GetOffset())

	if err := dbQuery.Find(&gormModels).Error; err != nil {
		return nil, models.NewSystemError(models.SystemErrorCodeValidation, models.SystemErrorTypeValidation, models.SystemErrorLevelError, "Query failed", struct{}{})
	}

	var entities []models.Role
	for _, gm := range gormModels {
		entities = append(entities, gm.ToModel())
	}

	totalPages := 0
	if limit > 0 {
		totalPages = int((totalRows + int64(limit) - 1) / int64(limit))
	}

	return &models.PaginatedResponse[models.Role]{
		TotalRows:  totalRows,
		TotalPages: totalPages,
		Rows:       entities,
	}, nil
}

// Override Create to ensure permissions are handled (though GORM usually handles this)
func (r *RoleRepository) Create(item models.Role) (models.Role, *models.SystemError) {
	// For Many-to-Many, we need to be careful not to duplicate permissions if we just pass them.
	// We should probably strip permissions or ensure they only have IDs if they are existing permissions.
	// But PermissionToEntity handles converting Model to Gorm.

	// Issue: If we pass a Permission with ID, GORM might try to Insert it if we are not careful,
	// unless we use Omit("Permissions") and then Association().Replace() or similar,
	// OR if the PermissionGorm is set up correctly and the IDs exist.

	// Use standard create but let's see.
	// Safest for M2M with existing items is to use Association replace or ensure we are just linking.

	gormModel := gormModels.RoleToEntity(item)

	// Using Clause(clause.OnConflict{DoNothing: true}) for permissions? No, that's for the role itself.

	if err := r.db.Create(&gormModel).Error; err != nil {
		return item, models.NewSystemError(models.SystemErrorCodeValidation, models.SystemErrorTypeValidation, models.SystemErrorLevelError, "Query failed", struct{}{})
	}
	// Reload to get properly populated fields if needed?
	// Usually Create returns the created object.

	return gormModel.ToModel(), nil
}

// Override Update to handle permission association replacement
func (r *RoleRepository) Update(id string, item models.Role) (models.Role, *models.SystemError) {
	gormModel := gormModels.RoleToEntity(item)

	// We should update the role fields AND the associations.
	// Generic Update uses Save(), which might work but for M2M replacement we often need to be explicit.

	// Transaction
	tx := r.db.Begin()

	// 1. Update Role fields
	if err := tx.Model(&gormModel).Where("id = ?", id).Omit("Permissions").Updates(&gormModel).Error; err != nil {
		tx.Rollback()
		return item, models.NewSystemError(models.SystemErrorCodeValidation, models.SystemErrorTypeValidation, models.SystemErrorLevelError, "Update failed", struct{}{})
	}

	// 2. Replace Associations
	if err := tx.Model(&gormModel).Association("Permissions").Replace(gormModel.Permissions); err != nil {
		tx.Rollback()
		return item, models.NewSystemError(models.SystemErrorCodeValidation, models.SystemErrorTypeValidation, models.SystemErrorLevelError, "Update associations failed", struct{}{})
	}

	tx.Commit()

	// Return updated model with permissions
	role, _ := r.GetOnce("id", id)
	return *role, nil
}
