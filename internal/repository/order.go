package repository

import (
	"context"
	models "rental-mobil/internal/model"

	"gorm.io/gorm"
)

type OrderRepo struct {
	DB *gorm.DB
}

func (r *OrderRepo) CreateOrder(ctx context.Context, order *models.Order) error {
	return r.DB.Create(order).Error
}

func (r *OrderRepo) FindByID(ctx context.Context, ID int) (models.Order, error) {
	var (
		resp = models.Order{}
	)
	if err := r.DB.Where("id = ?", ID).First(&resp).Error; err != nil {
		return resp, err
	}
	return resp, nil
}

func (r *OrderRepo) Update(ctx context.Context, req *models.Order) error {
	return r.DB.Save(req).Error
}

func (r *OrderRepo) Delete(ctx context.Context, ID int) error {
	return r.DB.Delete(&models.Order{}, ID).Error
}

func (r *OrderRepo) GetAll(ctx context.Context, objComp models.ComponentServerSide, param string) ([]models.Order, error) {
	var (
		resp []models.Order
	)
	if err := r.DB.Preload("Orders").Find(&resp).Error; err != nil {
		return resp, err
	}

	return resp, nil
}
