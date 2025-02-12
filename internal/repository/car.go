package repository

import (
	"context"
	models "rental-mobil/internal/model"

	"gorm.io/gorm"
)

type CarRepo struct {
	DB *gorm.DB
}

func (r *CarRepo) CreateCar(ctx context.Context, car *models.Car) error {
	return r.DB.Create(car).Error
}

func (r *CarRepo) FindByID(ctx context.Context, ID int) (models.Car, error) {
	var (
		resp = models.Car{}
	)
	if err := r.DB.Where("id = ?", ID).First(&resp).Error; err != nil {
		return resp, err
	}
	return resp, nil
}

func (r *CarRepo) Update(ctx context.Context, req *models.Car) error {
	return r.DB.Save(req).Error
}

func (r *CarRepo) Delete(ctx context.Context, ID int) error {
	return r.DB.Delete(&models.Car{}, ID).Error
}

func (r *CarRepo) GetAll(ctx context.Context, objComp models.ComponentServerSide, param string) ([]models.Car, error) {
	var (
		resp []models.Car
	)
	if err := r.DB.Preload("Orders").Find(&resp).Error; err != nil {
		return resp, err
	}

	return resp, nil
}
