package branddataservice

import (
	accessor "github.com/chiapingky/cars-go/internal/accessor-service"
)

type Brand struct {
	Id   int64
	Name string
}

type BrandAccessor interface {
	accessor.GenericAccessor[Brand, int64]
	FindByName(name string) (Brand, error)
}

type BrandService interface {
	GetBrandByName(name string) Brand
	InsertBrand(brand Brand) Brand
	DeleteBrand(brand Brand) Brand
}
