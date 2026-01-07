package postgress

import (
	"hrms/core/contracts"
	"hrms/core/models"
	"hrms/repository/postgress/repo"

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
	if err := db.AutoMigrate(&models.User{}, &models.Department{}); err != nil {
		return models.SystemError{
			Code:    models.SystemErrorCodeMigration,
			Type:    models.SystemErrorTypeValidation,
			Level:   models.SystemErrorLevelError,
			Message: "Failed to migrate database",
		}
	}
	return models.SystemError{}
}
