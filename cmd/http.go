package cmd

import (
	"log"
	"rental-mobil/helpers"
	"rental-mobil/internal/api"
	"rental-mobil/internal/interfaces"
	"rental-mobil/internal/repository"
	"rental-mobil/internal/services"

	"github.com/gin-gonic/gin"
)

func ServeHTTP() {
	d := dependencyInject()

	r := gin.Default()

	carV1 := r.Group("/car/v1")
	carV1.POST("/", d.CarAPI.Create)
	carV1.GET("/", d.CarAPI.GetAll)
	carV1.PUT("/:id", d.CarAPI.Update)
	carV1.DELETE("/:id", d.CarAPI.Delete)
	carV1.GET("/:id", d.CarAPI.Find)

	orderV1 := r.Group("/order/v1")
	orderV1.POST("/", d.OrderAPI.Create)
	orderV1.GET("/", d.OrderAPI.GetAll)
	orderV1.PUT("/:id", d.OrderAPI.Update)
	orderV1.DELETE("/:id", d.OrderAPI.Delete)
	orderV1.GET("/:id", d.OrderAPI.Find)

	err := r.Run(":" + helpers.GetEnv("PORT", ""))
	if err != nil {
		log.Fatal(err)
	}
}

type Dependency struct {
	CarAPI   interfaces.ICarAPI
	OrderAPI interfaces.IOrderAPI
}

func dependencyInject() Dependency {
	carRepo := &repository.CarRepo{
		DB: helpers.DB,
	}

	carSvc := &services.CarService{
		CarRepo: carRepo,
	}
	carAPI := &api.CarAPI{
		CarService: carSvc,
	}
	// order
	orderRepo := &repository.OrderRepo{
		DB: helpers.DB,
	}

	orderSvc := &services.OrderService{
		OrderRepo: orderRepo,
	}

	orderAPI := &api.OrderAPI{
		OrderService: orderSvc,
	}

	return Dependency{
		CarAPI:   carAPI,
		OrderAPI: orderAPI,
	}
}
