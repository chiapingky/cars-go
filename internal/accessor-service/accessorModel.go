package accessorservice

type GenericAccessor[T any, ID any] interface {
	FindAll() ([]T, error)
	FindById(id ID) (T, error)
	Save(data T) (T, error)
	Delete(data T) (T, error)
}
