package service

type BaseServiceInterface[T any] interface {
	GetAll() ([]T, error)
	GetByID(id uint) (T, error)
	Create(item T) (T, error)
	Update(item T) (T, error)
	Delete(id uint, item T) error
	HardDelete(id uint, item T) error
	Paginate(offset, limit int) ([]T, int64, error)
	Search(field, keyword string) ([]T, error)
	FindWithTrashed() ([]T, error)
	OnlyTrashed() ([]T, error)
	Restore(id uint, item T) error
}
