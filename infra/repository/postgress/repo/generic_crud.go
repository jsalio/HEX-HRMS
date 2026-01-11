package repo

import (
	"context"
	"sync"

	"hrms.local/core/models"

	"gorm.io/gorm"
)

type GenericCrud[T any, G any] struct {
	db       *gorm.DB
	mu       sync.RWMutex
	ctx      context.Context
	ToGorm   func(T) G
	ToEntity func(G) T
}

func NewGenericCrud[T any, G any](db *gorm.DB, toGorm func(T) G, toEntity func(G) T) GenericCrud[T, G] {
	return GenericCrud[T, G]{
		db:       db,
		ctx:      context.Background(),
		ToGorm:   toGorm,
		ToEntity: toEntity,
	}
}

// WithContext permite inyectar un context (normalmente desde un handler HTTP)
// Devuelve el mismo puntero para permitir encadenamiento
func (g *GenericCrud[T, G]) WithContext(ctx context.Context) *GenericCrud[T, G] {
	if ctx == nil {
		ctx = context.Background()
	}
	g.mu.Lock()
	g.ctx = ctx
	g.mu.Unlock()
	return g
}

// currentContext obtiene el contexto actual de forma segura
func (g *GenericCrud[T, G]) currentContext() context.Context {
	g.mu.RLock()
	defer g.mu.RUnlock()
	return g.ctx
}

func (g *GenericCrud[T, G]) GetByFilter(query models.SearchQuery) (*models.PaginatedResponse[T], *models.SystemError) {
	totalRows, sysErr := g.CountByFilter(query)
	if sysErr != nil {
		return nil, sysErr
	}

	var gormModels []G
	dbQuery := g.db.WithContext(g.currentContext())
	for _, filter := range query.Filters {
		dbQuery = dbQuery.Where(filter.Key+" = ?", filter.Value)
	}

	limit := query.Pagination.GetLimit()
	dbQuery = dbQuery.Limit(limit).Offset(query.Pagination.GetOffset())

	if err := dbQuery.Find(&gormModels).Error; err != nil {
		return nil, models.NewSystemError(models.SystemErrorCodeValidation, models.SystemErrorTypeValidation, models.SystemErrorLevelError, "Query failed", struct{}{})
	}

	var entities []T
	for _, gm := range gormModels {
		entities = append(entities, g.ToEntity(gm))
	}

	totalPages := 0
	if limit > 0 {
		totalPages = int((totalRows + int64(limit) - 1) / int64(limit))
	}

	return &models.PaginatedResponse[T]{
		TotalRows:  totalRows,
		TotalPages: totalPages,
		Rows:       entities,
	}, nil
}

func (g *GenericCrud[T, G]) CountByFilter(query models.SearchQuery) (int64, *models.SystemError) {
	var count int64
	var gormModel G
	dbQuery := g.db.WithContext(g.currentContext()).Model(&gormModel)
	for _, filter := range query.Filters {
		dbQuery = dbQuery.Where(filter.Key+" = ?", filter.Value)
	}

	if err := dbQuery.Count(&count).Error; err != nil {
		return 0, models.NewSystemError(models.SystemErrorCodeValidation, models.SystemErrorTypeValidation, models.SystemErrorLevelError, "Count failed", struct{}{})
	}

	return count, nil
}

func (g *GenericCrud[T, G]) Create(item T) (T, *models.SystemError) {
	gormModel := g.ToGorm(item)
	if err := g.db.WithContext(g.currentContext()).Create(&gormModel).Error; err != nil {
		return item, models.NewSystemError(models.SystemErrorCodeValidation, models.SystemErrorTypeValidation, models.SystemErrorLevelError, "Query failed", struct{}{})
	}
	return g.ToEntity(gormModel), nil
}

func (g *GenericCrud[T, G]) Update(id string, item T) (interface{}, *models.SystemError) {
	gormModel := g.ToGorm(item)
	if err := g.db.WithContext(g.currentContext()).Save(&gormModel).Error; err != nil {
		return nil, models.NewSystemError(models.SystemErrorCodeValidation, models.SystemErrorTypeValidation, models.SystemErrorLevelError, "Update failed", struct{}{})
	}
	return g.ToEntity(gormModel), nil
}

func (g *GenericCrud[T, G]) Delete(id string) (interface{}, error) {
	var gormModel G
	if err := g.db.WithContext(g.currentContext()).Where("id = ?", id).Delete(&gormModel).Error; err != nil {
		return nil, models.NewSystemError(models.SystemErrorCodeValidation, models.SystemErrorTypeValidation, models.SystemErrorLevelError, "Delete failed", struct{}{})
	}
	return nil, nil
}

func (g *GenericCrud[T, G]) GetOnce(key string, value any) (*T, *models.SystemError) {
	var gormModel G
	dbQuery := g.db.WithContext(g.currentContext()).Where(key+" = ?", value)
	if err := dbQuery.First(&gormModel).Error; err != nil {
		return nil, models.NewSystemError(models.SystemErrorCodeValidation, models.SystemErrorTypeValidation, models.SystemErrorLevelError, "GetOnce failed", struct{}{})
	}
	entity := g.ToEntity(gormModel)
	return &entity, nil
}

func (g *GenericCrud[T, G]) Exists(key string, value any) (bool, *models.SystemError) {
	var gormModel G
	dbQuery := g.db.WithContext(g.currentContext()).Where(key+" = ?", value)
	if err := dbQuery.First(&gormModel).Error; err != nil {
		return false, models.NewSystemError(models.SystemErrorCodeValidation, models.SystemErrorTypeValidation, models.SystemErrorLevelError, "Exists failed", struct{}{})
	}
	return true, nil
}
