package services

import (
	"context"
	"rental-mobil/internal/interfaces"
	models "rental-mobil/internal/model"
)

type OrderService struct {
	OrderRepo interfaces.IOrderRepo
}

func (s *OrderService) CreateOrder(ctx context.Context, order *models.Order) error {
	return s.OrderRepo.CreateOrder(ctx, order)
}

func (s *OrderService) FindByID(ctx context.Context, ID int) (models.Order, error) {
	return s.OrderRepo.FindByID(ctx, ID)
}

func (s *OrderService) Update(ctx context.Context, req *models.Order) error {
	// ambil data by id
	obj, err := s.FindByID(ctx, int(req.ID))
	if err != nil {
		return err
	}
	// perbarui data

	return s.OrderRepo.Update(ctx, &obj)
}

func (s *OrderService) Delete(ctx context.Context, ID int) error {
	return s.OrderRepo.Delete(ctx, ID)
}

func (s *OrderService) GetAll(ctx context.Context, objComp models.ComponentServerSide, param string) ([]models.Order, error) {
	return s.OrderRepo.GetAll(ctx, objComp, param)
}
