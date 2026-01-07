package repo

import (
	"context"
	"hrms/core/contracts"
	"hrms/core/models"
	"sync"

	"gorm.io/gorm"
)

type GenericCrud[T any] struct {
	db  *gorm.DB
	mu  sync.RWMutex    // protege el acceso concurrente al contexto
	ctx context.Context // contexto actual (por defecto Background)
	contracts.ReadOperation[T]
	contracts.WriteOperation[T]
}

func NewGenericCrud[T any](db *gorm.DB) GenericCrud[T] {
	return GenericCrud[T]{
		db:  db,
		ctx: context.Background(),
	}
}

// WithContext permite inyectar un context (normalmente desde un handler HTTP)
// Devuelve el mismo puntero para permitir encadenamiento
func (g *GenericCrud[T]) WithContext(ctx context.Context) *GenericCrud[T] {
	if ctx == nil {
		ctx = context.Background()
	}
	g.mu.Lock()
	g.ctx = ctx
	g.mu.Unlock()
	return g
}

// currentContext obtiene el contexto actual de forma segura
func (g *GenericCrud[T]) currentContext() context.Context {
	g.mu.RLock()
	defer g.mu.RUnlock()
	return g.ctx
}

func (g *GenericCrud[T]) GetByFilter(filters ...models.Filter) ([]T, *models.SystemError) {
	var gormModel []T
	query := g.db
	for _, filter := range filters {
		query = query.WithContext(g.currentContext()).Where(filter.Key+" = ?", filter.Value)
	}
	if err := query.Find(&gormModel).Error; err != nil {
		return nil, models.NewSystemError(models.SystemErrorCodeValidation, models.SystemErrorTypeValidation, models.SystemErrorLevelError, "Query failed", struct{}{})
	}
	return gormModel, nil
}

func (g *GenericCrud[T]) Create(item T) (T, *models.SystemError) {
	if err := g.db.WithContext(g.currentContext()).Create(&item).Error; err != nil {
		return item, models.NewSystemError(models.SystemErrorCodeValidation, models.SystemErrorTypeValidation, models.SystemErrorLevelError, "Query failed", struct{}{})
	}
	return item, nil
}

func (g *GenericCrud[T]) Update(id string, item T) (interface{}, *models.SystemError) {
	if err := g.db.WithContext(g.currentContext()).Save(&item).Error; err != nil {
		return nil, models.NewSystemError(models.SystemErrorCodeValidation, models.SystemErrorTypeValidation, models.SystemErrorLevelError, "Query failed", struct{}{})
	}
	return item, nil
}

func (g *GenericCrud[T]) Delete(id string) (interface{}, error) {
	if err := g.db.WithContext(g.currentContext()).Where("id = ?", id).Delete(new(T)).Error; err != nil {
		return nil, models.NewSystemError(models.SystemErrorCodeValidation, models.SystemErrorTypeValidation, models.SystemErrorLevelError, "Query failed", struct{}{})
	}
	return *new(T), nil
}
