package postgress

import (
	"strconv"
	"time"

	"github.com/google/uuid"
	"hrms.local/core/contracts"
	"hrms.local/core/models"
	gormModels "hrms.local/repository/postgress/models"
	"hrms.local/repository/postgress/repo"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Context struct {
	DB                 *gorm.DB
	UserContract       contracts.UserContract
	DepartmentContract contracts.DepartmentContract
	RoleContract       contracts.RoleContract
	PermissionContract contracts.PermissionContract
	PositionContract   contracts.PositionContract
}

func NewContext(dns string) (*Context, models.SystemError) {
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		return nil, models.SystemError{
			Code:    models.SystemErrorCodeValidation,
			Type:    models.SystemErrorTypeValidation,
			Level:   models.SystemErrorLevelError,
			Message: "Failed to connect to database",
		}
	}
	if err := migrate(db); err.Code != models.SystemErrorCodeNone {
		return nil, err
	}
	positionRepo := repo.NewPositionRepository(db)
	return &Context{
		DB:                 db,
		UserContract:       repo.NewUserRepository(db),
		DepartmentContract: repo.NewDepartmentRepository(db),
		RoleContract:       repo.NewRoleRepository(db),
		PermissionContract: repo.NewPermissionRepository(db),
		PositionContract:   &positionRepo,
	}, models.SystemError{}
}

func migrate(db *gorm.DB) models.SystemError {
	// Enable uuid-ossp extension
	if err := db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"").Error; err != nil {
		return models.SystemError{
			Code:    models.SystemErrorCodeMigration,
			Type:    models.SystemErrorTypeValidation,
			Level:   models.SystemErrorLevelError,
			Message: "Failed to create uuid-ossp extension",
		}
	}

	if err := db.AutoMigrate(
		&repo.UserGorm{},
		&repo.DepartmentGorm{},
		&repo.PositionGorm{},
		&gormModels.RoleGorm{},
		&gormModels.PermissionGorm{},
	); err != nil {
		return models.SystemError{
			Code:    models.SystemErrorCodeMigration,
			Type:    models.SystemErrorTypeValidation,
			Level:   models.SystemErrorLevelError,
			Message: "Failed to migrate database",
		}
	}
	if err := defaultUsers(db); err.Code != models.SystemErrorCodeNone {
		return err
	}
	if err := defaultPermissions(db); err.Code != models.SystemErrorCodeNone {
		return err
	}
	return models.SystemError{}
}

func defaultUsers(db *gorm.DB) models.SystemError {
	var gormModels []repo.UserGorm
	for i := range 100 {
		user := repo.UserGorm{
			ID:        uuid.New(),
			Username:  "user" + strconv.Itoa(i) + "@mail.com",
			Password:  "$2a$10$GjJPDCWa8Ig.7JC73mx6HuQ7yfsUblT5do8m3we9fUI3j34yaFlJ.", // password
			Name:      "User " + strconv.Itoa(i),
			Email:     "user" + strconv.Itoa(i) + "@mail.com",
			LastName:  "LastName " + strconv.Itoa(i),
			Picture:   "",
			Type:      "user",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			// DeletedAt: nil,
			Active: false,
			Role:   "user",
		}
		query := db.Where("email = ?", user.Email)
		query.Find(&gormModels)
		if len(gormModels) > 0 {
			continue
		}
		if err := query.Create(&user).Error; err != nil {
			return models.SystemError{
				Code:    models.SystemErrorCodeMigration,
				Type:    models.SystemErrorTypeValidation,
				Level:   models.SystemErrorLevelError,
				Message: "Failed to create default user",
			}
		}
	}

	return models.SystemError{}
}

func defaultPermissions(db *gorm.DB) models.SystemError {
	// Create default Admin role if not exists
	var adminRole gormModels.RoleGorm
	if err := db.Where("name = ?", "Admin").First(&adminRole).Error; err != nil {
		adminRole = gormModels.RoleGorm{
			ID:          uuid.New(),
			Name:        "Admin",
			Description: "Default Admin Role",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}
		if err := db.Create(&adminRole).Error; err != nil {
			return models.SystemError{
				Code:    models.SystemErrorCodeMigration,
				Type:    models.SystemErrorTypeValidation,
				Level:   models.SystemErrorLevelError,
				Message: "Failed to create default Admin role",
			}
		}
	}

	permissions := []models.Permission{
		{Name: models.PermissionViewMenuDashboard, Description: "View the dashboard menu"},
		{Name: models.PermissionViewMenuEmployees, Description: "View the employees menu"},
		{Name: models.PermissionEditEmployees, Description: "Edit employees"},
		{Name: models.PermissionViewEmployees, Description: "View employees"},
		{Name: models.PermissionViewMenuDepartments, Description: "View the departments menu"},
		{Name: models.PermissionViewMenuPosition, Description: "View the position menu"},
		{Name: models.PermissionViewMenuAttendance, Description: "View the attendance menu"},
		{Name: models.PermissionViewMenuPayroll, Description: "View the payroll menu"},
		{Name: models.PermissionViewMenuLeaveRequests, Description: "View the leave requests menu"},
		{Name: models.PermissionViewMenuSettings, Description: "View the settings menu"},
		{Name: models.PermissionAllAccess, Description: "Full access to settings"},
		{Name: models.PermissionViewRoles, Description: "View roles"},
		{Name: models.PermissionEditRoles, Description: "Edit roles"},
		{Name: models.PermissionEditUsers, Description: "Edit users"},
		{Name: models.PermissionViewUsers, Description: "View users"},
	}

	for _, p := range permissions {
		var existing gormModels.PermissionGorm
		if err := db.Where("name = ?", p.Name).First(&existing).Error; err == nil {
			continue
		}

		permGorm := gormModels.PermissionToEntity(p)
		if permGorm.ID == uuid.Nil {
			permGorm.ID = uuid.New()
		}
		permGorm.RoleID = adminRole.ID

		if err := db.Create(&permGorm).Error; err != nil {
			return models.SystemError{
				Code:    models.SystemErrorCodeMigration,
				Type:    models.SystemErrorTypeValidation,
				Level:   models.SystemErrorLevelError,
				Message: "Failed to create default permission: " + p.Name,
			}
		}
	}
	return models.SystemError{}
}
