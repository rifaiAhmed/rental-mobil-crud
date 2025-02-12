package interfaces

import (
	"context"
	models "rental-mobil/internal/model"

	"github.com/gin-gonic/gin"
)

type IOrderRepo interface {
	CreateOrder(ctx context.Context, order *models.Order) error
	FindByID(ctx context.Context, ID int) (models.Order, error)
	Update(ctx context.Context, req *models.Order) error
	Delete(ctx context.Context, ID int) error
	GetAll(ctx context.Context, objComp models.ComponentServerSide, param string) ([]models.Order, error)
}

type IOrderService interface {
	CreateOrder(ctx context.Context, order *models.Order) error
	FindByID(ctx context.Context, ID int) (models.Order, error)
	Update(ctx context.Context, req *models.Order) error
	Delete(ctx context.Context, ID int) error
	GetAll(ctx context.Context, objComp models.ComponentServerSide, param string) ([]models.Order, error)
}

type IOrderAPI interface {
	Create(c *gin.Context)
	GetAll(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	Find(c *gin.Context)
}
