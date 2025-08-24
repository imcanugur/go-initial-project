package repository

import (
	"context"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type BaseRepository[T any] struct {
	db *gorm.DB
}

func NewBaseRepository[T any](db *gorm.DB) *BaseRepository[T] {
	return &BaseRepository[T]{db}
}

// ---------------- BASIC CRUD ----------------

func (r *BaseRepository[T]) FindAll() ([]T, error) {
	var items []T
	err := r.db.Find(&items).Error
	return items, err
}

func (r *BaseRepository[T]) FindByID(id uint) (T, error) {
	var item T
	err := r.db.First(&item, id).Error
	return item, err
}

func (r *BaseRepository[T]) Create(item *T) error {
	return r.db.Create(item).Error
}

func (r *BaseRepository[T]) Update(item T) (T, error) {
	err := r.db.Save(&item).Error
	return item, err
}

func (r *BaseRepository[T]) Delete(id uint, item T) error {
	return r.db.Delete(&item, id).Error
}

func (r *BaseRepository[T]) HardDelete(id uint, item T) error {
	return r.db.Unscoped().Delete(&item, id).Error
}

// ---------------- BULK OPS ----------------

func (r *BaseRepository[T]) UpdateWhere(where map[string]interface{}, values map[string]interface{}) error {
	var item T
	return r.db.Model(&item).Where(where).Updates(values).Error
}

func (r *BaseRepository[T]) DeleteWhere(where map[string]interface{}) error {
	var item T
	return r.db.Where(where).Delete(&item).Error
}

func (r *BaseRepository[T]) CreateBatch(items []T, batchSize int) error {
	return r.db.CreateInBatches(items, batchSize).Error
}

// ---------------- FIND / FILTER ----------------

func (r *BaseRepository[T]) First(where map[string]interface{}) (T, error) {
	var item T
	err := r.db.Where(where).First(&item).Error
	return item, err
}

func (r *BaseRepository[T]) Where(where map[string]interface{}) ([]T, error) {
	var items []T
	err := r.db.Where(where).Find(&items).Error
	return items, err
}

func (r *BaseRepository[T]) Filter(where map[string]interface{}) ([]T, error) {
	var items []T
	query := r.db
	for key, val := range where {
		query = query.Where(key, val)
	}
	err := query.Find(&items).Error
	return items, err
}

func (r *BaseRepository[T]) Between(field string, from, to interface{}) ([]T, error) {
	var items []T
	err := r.db.Where(fmt.Sprintf("%s BETWEEN ? AND ?", field), from, to).Find(&items).Error
	return items, err
}

func (r *BaseRepository[T]) In(field string, values []interface{}) ([]T, error) {
	var items []T
	err := r.db.Where(fmt.Sprintf("%s IN ?", field), values).Find(&items).Error
	return items, err
}

func (r *BaseRepository[T]) NotIn(field string, values []interface{}) ([]T, error) {
	var items []T
	err := r.db.Where(fmt.Sprintf("%s NOT IN ?", field), values).Find(&items).Error
	return items, err
}

// ---------------- AGGREGATES ----------------

func (r *BaseRepository[T]) Count() (int64, error) {
	var count int64
	var item T
	err := r.db.Model(&item).Count(&count).Error
	return count, err
}

func (r *BaseRepository[T]) Sum(field string) (float64, error) {
	var result float64
	var item T
	err := r.db.Model(&item).Select("SUM(" + field + ")").Scan(&result).Error
	return result, err
}

func (r *BaseRepository[T]) Avg(field string) (float64, error) {
	var result float64
	var item T
	err := r.db.Model(&item).Select("AVG(" + field + ")").Scan(&result).Error
	return result, err
}

func (r *BaseRepository[T]) Min(field string) (float64, error) {
	var result float64
	var item T
	err := r.db.Model(&item).Select("MIN(" + field + ")").Scan(&result).Error
	return result, err
}

func (r *BaseRepository[T]) Max(field string) (float64, error) {
	var result float64
	var item T
	err := r.db.Model(&item).Select("MAX(" + field + ")").Scan(&result).Error
	return result, err
}

func (r *BaseRepository[T]) GroupBy(field string) ([]map[string]interface{}, error) {
	var results []map[string]interface{}
	var item T
	err := r.db.Model(&item).Select(field + ", COUNT(*) as count").Group(field).Scan(&results).Error
	return results, err
}

// ---------------- ORDER / PAGINATION ----------------

func (r *BaseRepository[T]) OrderBy(order string) ([]T, error) {
	var items []T
	err := r.db.Order(order).Find(&items).Error
	return items, err
}

func (r *BaseRepository[T]) OrderByMultiple(orders []string) ([]T, error) {
	var items []T
	query := r.db
	for _, order := range orders {
		query = query.Order(order)
	}
	err := query.Find(&items).Error
	return items, err
}

func (r *BaseRepository[T]) Paginate(offset int, limit int) ([]T, int64, error) {
	var items []T
	var count int64
	var item T
	query := r.db.Model(&item)
	query.Count(&count)
	err := query.Offset(offset).Limit(limit).Find(&items).Error
	return items, count, err
}

// ---------------- SEARCH ----------------

func (r *BaseRepository[T]) Search(field, keyword string) ([]T, error) {
	var items []T
	err := r.db.Where(field+" LIKE ?", "%"+keyword+"%").Find(&items).Error
	return items, err
}

// ---------------- SOFT DELETE ----------------

func (r *BaseRepository[T]) FindWithTrashed() ([]T, error) {
	var items []T
	err := r.db.Unscoped().Find(&items).Error
	return items, err
}

func (r *BaseRepository[T]) OnlyTrashed() ([]T, error) {
	var items []T
	err := r.db.Unscoped().Where("deleted_at IS NOT NULL").Find(&items).Error
	return items, err
}

func (r *BaseRepository[T]) Restore(id uint, item T) error {
	return r.db.Model(&item).Unscoped().Where("id = ?", id).Update("deleted_at", nil).Error
}

// ---------------- EXTRA POWER ----------------

// Join (raw join wrapper)
func (r *BaseRepository[T]) Join(query string, args ...interface{}) ([]T, error) {
	var items []T
	err := r.db.Joins(query, args...).Find(&items).Error
	return items, err
}

// Pluck tek alanı çek
func (r *BaseRepository[T]) Pluck(field string) ([]interface{}, error) {
	var results []interface{}
	var item T
	err := r.db.Model(&item).Pluck(field, &results).Error
	return results, err
}

// Chunk – büyük dataset’i parça parça işleme
func (r *BaseRepository[T]) Chunk(size int, fn func([]T) error) error {
	var items []T
	tx := r.db
	for {
		result := tx.Limit(size).Find(&items)
		if result.Error != nil {
			return result.Error
		}
		if len(items) == 0 {
			break
		}
		if err := fn(items); err != nil {
			return err
		}
		if int(result.RowsAffected) < size {
			break
		}
	}
	return nil
}

// DebugSQL Debug – son SQL
func (r *BaseRepository[T]) DebugSQL() *gorm.DB {
	return r.db.Debug()
}

// Context destekli
func (r *BaseRepository[T]) FindAllCtx(ctx context.Context) ([]T, error) {
	var items []T
	err := r.db.WithContext(ctx).Find(&items).Error
	return items, err
}

// Transaction içinde repo instance
func (r *BaseRepository[T]) WithTransactionRepo(fn func(repo BaseRepositoryInterface[T]) error) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		subRepo := NewBaseRepository[T](tx)
		return fn(subRepo)
	})
}

// Upsert
func (r *BaseRepository[T]) Upsert(item T, conflictColumns []string) error {
	return r.db.Clauses(clause.OnConflict{
		Columns:   toClauseColumns(conflictColumns),
		UpdateAll: true,
	}).Create(&item).Error
}

// Helper
func toClauseColumns(cols []string) []clause.Column {
	result := make([]clause.Column, len(cols))
	for i, c := range cols {
		result[i] = clause.Column{Name: c}
	}
	return result
}

// ---------------- BUL / FİLTRELE ----------------

// Field seçerek getir
func (r *BaseRepository[T]) Select(fields []string) ([]T, error) {
	var items []T
	err := r.db.Select(fields).Find(&items).Error
	return items, err
}

// Belirli şartlarla ilk kaydı getir veya oluştur
func (r *BaseRepository[T]) FirstOrCreate(where map[string]interface{}, defaults T) (T, error) {
	var item T
	err := r.db.Where(where).FirstOrCreate(&item, defaults).Error
	return item, err
}

func (r *BaseRepository[T]) Exists(where map[string]interface{}) (bool, error) {
	var item T
	err := r.db.Where(where).First(&item).Error
	if err == gorm.ErrRecordNotFound {
		return false, nil
	}
	return err == nil, err
}

// ---------------- UPDATE / UPSERT ----------------
func (r *BaseRepository[T]) UpdateColumns(id uint, values map[string]interface{}) error {
	var item T
	return r.db.Model(&item).Where("id = ?", id).Updates(values).Error
}

// Distinct field değerleri
func (r *BaseRepository[T]) Distinct(field string) ([]interface{}, error) {
	var results []interface{}
	var item T
	err := r.db.Model(&item).Distinct(field).Pluck(field, &results).Error
	return results, err
}

// Scope destekli sorgu
func (r *BaseRepository[T]) WithScopes(scopes ...func(*gorm.DB) *gorm.DB) ([]T, error) {
	var items []T
	err := r.db.Scopes(scopes...).Find(&items).Error
	return items, err
}

// Preload ile ilişkili veriler
func (r *BaseRepository[T]) WithPreload(preloads []string) ([]T, error) {
	var items []T
	query := r.db
	for _, preload := range preloads {
		query = query.Preload(preload)
	}
	err := query.Find(&items).Error
	return items, err
}

// Raw SQL
func (r *BaseRepository[T]) RawQuery(sql string, values ...interface{}) (*gorm.DB, error) {
	tx := r.db.Raw(sql, values...)
	return tx, tx.Error
}

// Transaction
func (r *BaseRepository[T]) Transaction(fn func(tx *gorm.DB) error) error {
	return r.db.Transaction(fn)
}

// Lock for update
func (r *BaseRepository[T]) FindForUpdate(id uint) (T, error) {
	var item T
	err := r.db.Clauses(clause.Locking{Strength: "UPDATE"}).First(&item, id).Error
	return item, err
}

// ---------------- CONTEXT DESTEKLİ ----------------

func (r *BaseRepository[T]) FindByIDCtx(ctx context.Context, id uint) (T, error) {
	var item T
	err := r.db.WithContext(ctx).First(&item, id).Error
	return item, err
}
