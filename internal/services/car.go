package services

import (
	"context"
	"rental-mobil/internal/interfaces"
	models "rental-mobil/internal/model"
)

type CarService struct {
	CarRepo interfaces.ICarRepo
}

func (s *CarService) CreateCar(ctx context.Context, car *models.Car) error {
	return s.CarRepo.CreateCar(ctx, car)
}

func (s *CarService) FindByID(ctx context.Context, ID int) (models.Car, error) {
	return s.CarRepo.FindByID(ctx, ID)
}

func (s *CarService) Update(ctx context.Context, req *models.Car) error {
	// ambil data by id
	obj, err := s.FindByID(ctx, int(req.ID))
	if err != nil {
		return err
	}
	// perbarui data
	obj.CarName = req.CarName
	obj.DayRate = req.DayRate
	obj.MonthRate = req.MonthRate
	obj.ImageCar = req.ImageCar

	return s.CarRepo.Update(ctx, &obj)
}

func (s *CarService) Delete(ctx context.Context, ID int) error {
	return s.CarRepo.Delete(ctx, ID)
}

func (s *CarService) GetAll(ctx context.Context, objComp models.ComponentServerSide, param string) ([]models.Car, error) {
	return s.CarRepo.GetAll(ctx, objComp, param)
}
