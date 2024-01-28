package branddataservice

import (
	"database/sql"
	"errors"
	"log"
)

type BrandAccessorImpl struct {
	DB *sql.DB
}

func (accessor BrandAccessorImpl) FindAll() ([]Brand, error) {
	rows, err := accessor.DB.Query("SELECT Id, Name FROM brand")
	if err != nil {
		return make([]Brand, 0), errors.New("FindAll() | Error when querying brand " + err.Error())
	}
	var result []Brand
	for rows.Next() {
		var temp Brand
		err = rows.Scan(&temp.Id, &temp.Name)
		if err != nil {
			log.Default().Println("FindAll() | An entry can't be loaded ", err)
		} else {
			result = append(result, temp)
		}
	}
	return result, nil
}

func (accessor BrandAccessorImpl) FindById(id int64) (Brand, error) {
	row := accessor.DB.QueryRow("SELECT Id, Name FROM brand WHERE Id = ?", id)
	var result Brand
	err := row.Scan(&result.Id, &result.Name)
	if err != nil {
		return Brand{}, errors.New("FindById | Error when querying brand " + err.Error())
	}
	return result, nil
}

func (accessor BrandAccessorImpl) Save(data Brand) (Brand, error) {
	row := accessor.DB.QueryRow("INSERT INTO brand (Name) VALUES (?) RETURNING Id", data.Name)
	var id int64
	err := row.Scan(&id)
	if err != nil {
		return Brand{}, errors.New("Save | Error when saving brand " + err.Error())
	}
	result, err := accessor.FindById(id)
	if err != nil {
		return Brand{}, errors.New("Save | Failed saving brand " + err.Error())
	}
	return result, nil
}

func (accessor BrandAccessorImpl) Delete(data Brand) (Brand, error) {
	row := accessor.DB.QueryRow("DELETE FROM brand WHERE Name = ? RETURNING Id", data.Name)
	var id int64
	err := row.Scan(&id)
	if err != nil {
		return Brand{}, errors.New("Delete | Error when deleting brand " + err.Error())
	}
	result, err := accessor.FindById(id)
	if err == nil {
		return result, errors.New("Delete | Failed deleting brand " + err.Error())
	}
	data.Id = id
	return data, nil
}

func (accessor BrandAccessorImpl) FindByName(name string) (Brand, error) {
	row := accessor.DB.QueryRow("SELECT Id, Name FROM brand WHERE Name = ?", name)
	var result Brand
	err := row.Scan(&result.Id, &result.Name)
	if err != nil {
		return Brand{}, errors.New("FindById | Error when querying brand " + err.Error())
	}
	return result, nil
}
