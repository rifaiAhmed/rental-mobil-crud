package interfaces

import (
	"context"
	models "rental-mobil/internal/model"

	"github.com/gin-gonic/gin"
)

type ICarRepo interface {
	CreateCar(ctx context.Context, car *models.Car) error
	FindByID(ctx context.Context, ID int) (models.Car, error)
	Update(ctx context.Context, req *models.Car) error
	Delete(ctx context.Context, ID int) error
	GetAll(ctx context.Context, objComp models.ComponentServerSide, param string) ([]models.Car, error)
}

type ICarService interface {
	CreateCar(ctx context.Context, car *models.Car) error
	FindByID(ctx context.Context, ID int) (models.Car, error)
	Update(ctx context.Context, req *models.Car) error
	Delete(ctx context.Context, ID int) error
	GetAll(ctx context.Context, objComp models.ComponentServerSide, param string) ([]models.Car, error)
}

type ICarAPI interface {
	Create(c *gin.Context)
	GetAll(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	Find(c *gin.Context)
}
