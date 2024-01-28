package requestcontrollerservice

import (
	"log"
	"net/http"

	car "github.com/chiapingky/cars-go/internal/car-data-service"

	brand "github.com/chiapingky/cars-go/internal/brand-data-service"

	"github.com/gin-gonic/gin"
)

type ControllerService struct {
	DelegateService
}

func (controller *ControllerService) GetAllCarsOrCarsByBrand(c *gin.Context) {
	brand, isExist := c.GetQuery("brand")
	if !isExist {
		cars := controller.DelegateService.GetAllCars()
		c.JSON(http.StatusOK, cars)
		return
	} else if len([]rune(brand)) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Car must not be empty",
		})
		return
	}
	cars := controller.DelegateService.GetCarByBrandName(brand)
	c.JSON(http.StatusOK, cars)
}

func (controller *ControllerService) InsertCar(c *gin.Context) {
	var carData car.Car = car.Car{}
	err := c.BindJSON(&carData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
		})
		return
	}
	if carData.Brand.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Please insert valid brand",
		})
		return
	}
	if carData.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Please insert valid car",
		})
		return
	}
	brandData := controller.DelegateService.InsertBrand(carData.Brand)
	if (brandData == brand.Brand{}) {
		log.Default().Println("Brand does not exist, new brand inserted")
		brandData = controller.DelegateService.GetBrandByName(brandData.Name)
	}
	carData.Brand = brandData
	result := controller.DelegateService.InsertCar(carData)
	c.JSON(http.StatusOK, result)
}

func (controller *ControllerService) DeleteCar(c *gin.Context) {
	/*
		var input map[string]string
		err := c.BindJSON(&input)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid request",
			})
			return
		}
		brandName := input["brandName"]
		carName := input["carName"]
		if brandName == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Please insert valid brand",
			})
			return
		}
		if carName == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Please insert valid car",
			})
			return
		}
		result := carservice.DeleteCar(carName, brandName)
		c.JSON(http.StatusOK, result)
	*/
}

func (controller *ControllerService) InsertBrand(c *gin.Context) {
	/*
		var brand string
		err := c.BindJSON(&brand)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid request",
			})
			return
		}
		result := brandservice.InsertBrand(brand)
		c.JSON(http.StatusOK, result)
	*/
}

func (controller *ControllerService) DeleteBrand(c *gin.Context) {
	/*
		var brand string
		err := c.BindJSON(&brand)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid request",
			})
			return
		}
		brandData := brandservice.GetBrandByName(brand)
		if brandData == "" {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Brand does not exist",
			})
			return
		}
		result := brandservice.DeleteBrand(brand)
		c.JSON(http.StatusOK, result)
	*/
}
