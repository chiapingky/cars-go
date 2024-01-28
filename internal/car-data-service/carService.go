package cardataservice

import "log"

type CarServiceImpl struct {
	CarAccessor
}

func (carService CarServiceImpl) GetAllCars() []Car {
	cars, err := carService.CarAccessor.FindAll()
	if err != nil {
		log.Default().Println("GetAllCars() | Error when getting car ", err)
		return make([]Car, 0)
	}
	return cars
}

func (carService CarServiceImpl) GetCarByBrandName(brandName string) []Car {
	cars, err := carService.CarAccessor.FindByBrandName(brandName)
	if err != nil {
		log.Default().Println("GetCarByBrandName() | Error when getting car ", err)
		return make([]Car, 0)
	}
	return cars
}

func (carService CarServiceImpl) InsertCar(car Car) Car {
	car, err := carService.FindByNameAndBrandId(car.Name, car.Brand.Id)
	if err == nil {
		log.Default().Println("InsertCar() | Car already exist ")
		return Car{}
	}
	saveResult, err := carService.CarAccessor.Save(car)
	if err != nil {
		log.Default().Println("InsertCar() | Error when saving Car ", err)
		return Car{}
	}
	return saveResult
}

func (carService CarServiceImpl) DeleteCar(car Car) Car {
	car, err := carService.FindByNameAndBrandId(car.Name, car.Brand.Id)
	if err != nil {
		log.Default().Println("DeleteCar() | Car doesn't exist ")
		return Car{}
	}
	deleteResult, err := carService.CarAccessor.Delete(car)
	if err != nil {
		log.Default().Println("DeleteCar() | Error when deleting car ", err)
		return Car{}
	}
	return deleteResult
}
