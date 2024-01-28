package cardataservice

import (
	accessor "github.com/chiapingky/cars-go/internal/accessor-service"

	brand "github.com/chiapingky/cars-go/internal/brand-data-service"
)

type Car struct {
	Id    int32
	Name  string
	Brand brand.Brand
}

type CarAccessor interface {
	accessor.GenericAccessor[Car, int32]
	FindByBrandName(brandName string) ([]Car, error)
	FindByNameAndBrandId(carName string, brandId int64) (Car, error)
}

type CarService interface {
	GetAllCars() []Car
	GetCarByBrandName(brandName string) []Car
	InsertCar(car Car) Car
	DeleteCar(car Car) Car
}
