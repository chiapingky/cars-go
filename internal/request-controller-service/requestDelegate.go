package requestcontrollerservice

import (
	car "github.com/chiapingky/cars-go/internal/car-data-service"

	brand "github.com/chiapingky/cars-go/internal/brand-data-service"
)

type DelegateService struct {
	brand.BrandService
	car.CarService
}

func (delegate *DelegateService) GetAllCars() []car.Car {
	return delegate.CarService.GetAllCars()
}

func (delegate *DelegateService) GetCarByBrand(brand string) []car.Car {
	return delegate.CarService.GetCarByBrandName(brand)
}

func (delegate *DelegateService) InsertBrand(brandName brand.Brand) brand.Brand {
	return delegate.BrandService.InsertBrand(brandName)
}

func (delegate *DelegateService) DeleteBrand(brandName brand.Brand) brand.Brand {
	return delegate.BrandService.DeleteBrand(brandName)
}

func (delegate *DelegateService) InsertCar(carData car.Car) car.Car {
	return delegate.CarService.InsertCar(carData)
}

func (delegate *DelegateService) DeleteCar(carData car.Car) car.Car {
	return delegate.CarService.DeleteCar(carData)
}
