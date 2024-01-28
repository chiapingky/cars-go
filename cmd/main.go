package main

import (
	"database/sql"
	"fmt"
	"log"

	controller "github.com/chiapingky/cars-go/internal/request-controller-service"

	car "github.com/chiapingky/cars-go/internal/car-data-service"

	brand "github.com/chiapingky/cars-go/internal/brand-data-service"

	"github.com/gin-gonic/gin"
)

func main() {
	printLogo()

	var db *sql.DB
	db, err := sql.Open("postgres", "postgres://username:password@localhost:5432/db?sslmode=disabled")
	if err != nil {
		log.Fatal(err)
	}
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected to db")

	brandAccessor := brand.BrandAccessorImpl{DB: db}
	carAccessor := car.CarAccessorImpl{DB: db}

	brandService := brand.BrandServiceImpl{BrandAccessor: brandAccessor}
	carService := car.CarServiceImpl{CarAccessor: carAccessor}

	delegate := controller.DelegateService{BrandService: brandService, CarService: carService}
	controller := controller.ControllerService{DelegateService: delegate}

	router := gin.Default()

	router.GET("/car", controller.GetAllCarsOrCarsByBrand)
	router.POST("/car", controller.InsertCar)
	router.DELETE("/car", controller.DeleteCar)
	router.POST("/car/brand", controller.InsertBrand)
	router.DELETE("/brand", controller.DeleteBrand)

	log.Default().Println("Service started")
	router.Run("localhost:8080")
}

func printLogo() {
	logo := `
	  _____          _____                 _          
	 / ____|        / ____|               (_)         
	| |  __  ___   | (___   ___ _ ____   ___  ___ ___ 
	| | |_ |/ _ \   \___ \ / _ \ '__\ \ / / |/ __/ _ \
	| |__| | (_) |  ____) |  __/ |   \ V /| | (_|  __/
	 \_____|\___/  |_____/ \___|_|    \_/ |_|\___\___|
`
	fmt.Println(logo)
}
