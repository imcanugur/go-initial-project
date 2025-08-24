package service

import (
	"context"
	"go-initial-project/repository"
)

type BaseService[T any] struct {
	repo repository.BaseRepositoryInterface[T]
}

func NewBaseService[T any](repo repository.BaseRepositoryInterface[T]) *BaseService[T] {
	return &BaseService[T]{repo: repo}
}

// ---------------- BASIC CRUD ----------------
func (s *BaseService[T]) GetAll() ([]T, error)       { return s.repo.FindAll() }
func (s *BaseService[T]) GetByID(id uint) (T, error) { return s.repo.FindByID(id) }

func (s *BaseService[T]) Create(item T) (T, error) {
	err := s.repo.Create(&item) // &item → pointer
	return item, err
}

func (s *BaseService[T]) Update(item T) (T, error)         { return s.repo.Update(item) }
func (s *BaseService[T]) Delete(id uint, item T) error     { return s.repo.Delete(id, item) }
func (s *BaseService[T]) HardDelete(id uint, item T) error { return s.repo.HardDelete(id, item) }

// ---------------- BULK ----------------
func (s *BaseService[T]) UpdateWhere(where map[string]interface{}, values map[string]interface{}) error {
	return s.repo.UpdateWhere(where, values)
}
func (s *BaseService[T]) DeleteWhere(where map[string]interface{}) error {
	return s.repo.DeleteWhere(where)
}
func (s *BaseService[T]) CreateBatch(items []T, batchSize int) error {
	return s.repo.CreateBatch(items, batchSize)
}

// ---------------- FIND / FILTER ----------------
func (s *BaseService[T]) First(where map[string]interface{}) (T, error) {
	return s.repo.First(where)
}
func (s *BaseService[T]) Where(where map[string]interface{}) ([]T, error) {
	return s.repo.Where(where)
}
func (s *BaseService[T]) Filter(where map[string]interface{}) ([]T, error) {
	return s.repo.Filter(where)
}
func (s *BaseService[T]) Between(field string, from, to interface{}) ([]T, error) {
	return s.repo.Between(field, from, to)
}
func (s *BaseService[T]) In(field string, values []interface{}) ([]T, error) {
	return s.repo.In(field, values)
}
func (s *BaseService[T]) NotIn(field string, values []interface{}) ([]T, error) {
	return s.repo.NotIn(field, values)
}

// ---------------- AGGREGATES ----------------
func (s *BaseService[T]) Count() (int64, error)             { return s.repo.Count() }
func (s *BaseService[T]) Sum(field string) (float64, error) { return s.repo.Sum(field) }
func (s *BaseService[T]) Avg(field string) (float64, error) { return s.repo.Avg(field) }
func (s *BaseService[T]) Min(field string) (float64, error) { return s.repo.Min(field) }
func (s *BaseService[T]) Max(field string) (float64, error) { return s.repo.Max(field) }
func (s *BaseService[T]) GroupBy(field string) ([]map[string]interface{}, error) {
	return s.repo.GroupBy(field)
}

// ---------------- ORDER & PAGINATION ----------------
func (s *BaseService[T]) OrderBy(order string) ([]T, error) {
	return s.repo.OrderBy(order)
}
func (s *BaseService[T]) OrderByMultiple(orders []string) ([]T, error) {
	return s.repo.OrderByMultiple(orders)
}
func (s *BaseService[T]) Paginate(offset int, limit int) ([]T, int64, error) {
	return s.repo.Paginate(offset, limit)
}

// ---------------- SEARCH ----------------
func (s *BaseService[T]) Search(field, keyword string) ([]T, error) {
	return s.repo.Search(field, keyword)
}

// ---------------- SOFT DELETE ----------------
func (s *BaseService[T]) FindWithTrashed() ([]T, error) {
	return s.repo.FindWithTrashed()
}
func (s *BaseService[T]) OnlyTrashed() ([]T, error) {
	return s.repo.OnlyTrashed()
}
func (s *BaseService[T]) Restore(id uint, item T) error {
	return s.repo.Restore(id, item)
}

// ---------------- EXTRA ----------------
func (s *BaseService[T]) Join(query string, args ...interface{}) ([]T, error) {
	return s.repo.Join(query, args...)
}
func (s *BaseService[T]) Pluck(field string) ([]interface{}, error) {
	return s.repo.Pluck(field)
}
func (s *BaseService[T]) Chunk(size int, fn func([]T) error) error {
	return s.repo.Chunk(size, fn)
}
func (s *BaseService[T]) DebugSQL() any {
	return s.repo.DebugSQL()
}

// ---------------- CONTEXT & TX ----------------
func (s *BaseService[T]) GetAllCtx(ctx context.Context) ([]T, error) {
	return s.repo.FindAllCtx(ctx)
}
func (s *BaseService[T]) GetByIDCtx(ctx context.Context, id uint) (T, error) {
	return s.repo.FindByIDCtx(ctx, id)
}
func (s *BaseService[T]) WithTransaction(fn func(repo repository.BaseRepositoryInterface[T]) error) error {
	return s.repo.WithTransactionRepo(fn)
}

// ---------------- UPSERT & LOCK ----------------
func (s *BaseService[T]) Upsert(item T, conflictColumns []string) error {
	return s.repo.Upsert(item, conflictColumns)
}
func (s *BaseService[T]) FindForUpdate(id uint) (T, error) {
	return s.repo.FindForUpdate(id)
}
