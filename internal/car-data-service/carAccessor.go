package cardataservice

import (
	"database/sql"
	"errors"
	"log"
)

var JOIN_QUERY string = "SELECT car.Id AS Id, car.Name AS Name, brand.Id AS BrandId, brand.Name AS BrandName FROM car LEFT JOIN brand ON car.BrandId = brand.Id"

type CarAccessorImpl struct {
	DB *sql.DB
}

func (accessor CarAccessorImpl) FindAll() ([]Car, error) {
	rows, err := accessor.DB.Query("SELECT Id, Name, BrandId, BrandName FROM (?)", JOIN_QUERY)
	if err != nil {
		return make([]Car, 0), errors.New("FindAll() | Error when querying car " + err.Error())
	}
	var result []Car
	for rows.Next() {
		var temp Car
		err = rows.Scan(&temp.Id, &temp.Name, &temp.Brand.Id, &temp.Brand.Name)
		if err != nil {
			log.Default().Println("FindAll() | An entry can't be loaded ", err)
		} else {
			result = append(result, temp)
		}
	}
	return result, nil
}

func (accessor CarAccessorImpl) FindById(id int32) (Car, error) {
	row := accessor.DB.QueryRow("SELECT Id, Name, BrandId, BrandName FROM ((?) AS d) WHERE Id = ?", JOIN_QUERY, id)
	var result Car
	err := row.Scan(&result.Id, &result.Name, &result.Brand.Id, &result.Brand.Name)
	if err != nil {
		return Car{}, errors.New("FindById | Error when querying car " + err.Error())
	}
	return result, nil
}

func (accessor CarAccessorImpl) Save(data Car) (Car, error) {
	row := accessor.DB.QueryRow("INSERT INTO car (Name) VALUES (?) RETURNING Id", data.Name)
	var id int32
	err := row.Scan(&id)
	if err != nil {
		return Car{}, errors.New("Save | Error when saving car " + err.Error())
	}
	result, err := accessor.FindById(id)
	if err != nil {
		return Car{}, errors.New("Save | Failed saving car " + err.Error())
	}
	return result, nil
}

func (accessor CarAccessorImpl) Delete(data Car) (Car, error) {
	row := accessor.DB.QueryRow("DELETE FROM car WHERE Name = ? RETURNING Id", data.Name)
	var id int32
	err := row.Scan(&id)
	if err != nil {
		return Car{}, errors.New("Delete | Error when deleting car " + err.Error())
	}
	result, err := accessor.FindById(id)
	if err == nil {
		return result, errors.New("Delete | Failed deleting car " + err.Error())
	}
	data.Id = id
	return data, nil
}

func (accessor CarAccessorImpl) FindByBrandName(brandName string) ([]Car, error) {
	rows, err := accessor.DB.Query("SELECT Id, Name, BrandId, BrandName FROM ((?) AS d) WHERE BrandName = ?", JOIN_QUERY, brandName)
	if err != nil {
		return make([]Car, 0), errors.New("FindByBrandName() | Error when querying car " + err.Error())
	}
	var result []Car
	for rows.Next() {
		var temp Car
		err = rows.Scan(&temp.Id, &temp.Name)
		if err != nil {
			log.Default().Println("FindByBrandName() | An entry can't be loaded ", err)
		} else {
			result = append(result, temp)
		}
	}
	return result, nil
}

func (accessor CarAccessorImpl) FindByNameAndBrandId(carName string, brandId int64) (Car, error) {
	row := accessor.DB.QueryRow("SELECT Id, Name, BrandId, BrandName FROM ((?) AS d) WHERE Name = ? AND  BrandId = ?", JOIN_QUERY, carName, brandId)
	var result Car
	err := row.Scan(&result.Id, &result.Name, &result.Brand.Id, &result.Brand.Name)
	if err != nil {
		return Car{}, errors.New("FindById | Error when querying car " + err.Error())
	}
	return result, nil
}
