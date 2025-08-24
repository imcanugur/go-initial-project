package repository

import (
	"context"

	"gorm.io/gorm"
)

type BaseRepositoryInterface[T any] interface {
	FindAll() ([]T, error)
	FindByID(id uint) (T, error)
	Create(item *T) error
	Update(item T) (T, error)
	Delete(id uint, item T) error
	HardDelete(id uint, item T) error

	UpdateWhere(where map[string]interface{}, values map[string]interface{}) error
	DeleteWhere(where map[string]interface{}) error
	CreateBatch(items []T, batchSize int) error

	First(where map[string]interface{}) (T, error)
	Where(where map[string]interface{}) ([]T, error)
	Filter(where map[string]interface{}) ([]T, error)
	Between(field string, from, to interface{}) ([]T, error)
	In(field string, values []interface{}) ([]T, error)
	NotIn(field string, values []interface{}) ([]T, error)

	Count() (int64, error)
	Sum(field string) (float64, error)
	Avg(field string) (float64, error)
	Min(field string) (float64, error)
	Max(field string) (float64, error)
	GroupBy(field string) ([]map[string]interface{}, error)

	OrderBy(order string) ([]T, error)
	OrderByMultiple(orders []string) ([]T, error)
	Paginate(offset int, limit int) ([]T, int64, error)

	Search(field, keyword string) ([]T, error)

	FindWithTrashed() ([]T, error)
	OnlyTrashed() ([]T, error)
	Restore(id uint, item T) error

	Join(query string, args ...interface{}) ([]T, error)
	Pluck(field string) ([]interface{}, error)
	Chunk(size int, fn func([]T) error) error
	DebugSQL() *gorm.DB

	FindAllCtx(ctx context.Context) ([]T, error)
	FindByIDCtx(ctx context.Context, id uint) (T, error)

	WithTransactionRepo(fn func(repo BaseRepositoryInterface[T]) error) error

	Upsert(item T, conflictColumns []string) error
	FindForUpdate(id uint) (T, error)
}
