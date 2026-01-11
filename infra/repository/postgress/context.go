package postgress

import (
	"strconv"
	"time"

	"github.com/google/uuid"
	"hrms.local/core/contracts"
	"hrms.local/core/models"
	"hrms.local/repository/postgress/repo"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Context struct {
	DB                 *gorm.DB
	UserContract       contracts.UserContract
	DepartmentContract contracts.DepartmentContract
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
	return &Context{
		DB:                 db,
		UserContract:       repo.NewUserRepository(db),
		DepartmentContract: repo.NewDepartmentRepository(db),
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

	if err := db.AutoMigrate(&repo.UserGorm{}, &repo.DepartmentGorm{}); err != nil {
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

	// user := repo.UserGorm{
	// 	ID:       uuid.New(),
	// 	Username: "admin",
	// 	Password: "admin",
	// 	Role:     "admin",
	// }
	// if err := db.Create(&user).Error; err != nil {
	// 	return models.SystemError{
	// 		Code:    models.SystemErrorCodeMigration,
	// 		Type:    models.SystemErrorTypeValidation,
	// 		Level:   models.SystemErrorLevelError,
	// 		Message: "Failed to create default user",
	// 	}
	// }
	// return models.SystemError{}
}
