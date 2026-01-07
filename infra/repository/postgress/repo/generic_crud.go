package repo

import (
	"context"
	"hrms/core/contracts"
	"hrms/core/models"

	"gorm.io/gorm"
)

type GenericCrud[T any] struct {
	db *gorm.DB
	contracts.ReadOperation[T]
	contracts.WriteOperation[T]
}

func NewGenericCrud[T any](db *gorm.DB) GenericCrud[T] {
	return GenericCrud[T]{db: db}
}

func (g *GenericCrud[T]) GetByFilter(filters ...models.Filter) ([]T, *models.SystemError) {
	var gormModel []T
	query := g.db
	for _, filter := range filters {
		query = query.WithContext(context.Background()).Where(filter.Key+" = ?", filter.Value)
	}
	if err := query.Find(&gormModel).Error; err != nil {
		return nil, models.NewSystemError(models.SystemErrorCodeValidation, models.SystemErrorTypeValidation, models.SystemErrorLevelError, "Query failed", struct{}{})
	}
	return gormModel, nil
}

func (g *GenericCrud[T]) Create(item T) (T, *models.SystemError) {
	if err := g.db.Create(&item).Error; err != nil {
		return item, models.NewSystemError(models.SystemErrorCodeValidation, models.SystemErrorTypeValidation, models.SystemErrorLevelError, "Query failed", struct{}{})
	}
	return item, nil
}

func (g *GenericCrud[T]) Update(id string, item T) (interface{}, *models.SystemError) {
	if err := g.db.Save(&item).Error; err != nil {
		return nil, models.NewSystemError(models.SystemErrorCodeValidation, models.SystemErrorTypeValidation, models.SystemErrorLevelError, "Query failed", struct{}{})
	}
	return item, nil
}

func (g *GenericCrud[T]) Delete(id string) (interface{}, error) {
	if err := g.db.Where("id = ?", id).Delete(new(T)).Error; err != nil {
		return nil, models.NewSystemError(models.SystemErrorCodeValidation, models.SystemErrorTypeValidation, models.SystemErrorLevelError, "Query failed", struct{}{})
	}
	return *new(T), nil
}
