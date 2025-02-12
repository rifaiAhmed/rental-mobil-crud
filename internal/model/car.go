package models

import "github.com/go-playground/validator/v10"

type Car struct {
	ID        uint    `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	CarName   string  `json:"car_name" gorm:"type:varchar(255);not null"`
	DayRate   float64 `json:"day_rate" gorm:"type:double;not null"`
	MonthRate float64 `json:"month_rate" gorm:"type:double;not null"`
	ImageCar  string  `json:"image_car" gorm:"type:varchar(255);not null"`
	Orders    []Order `json:"orders" gorm:"foreignKey:CarID"`
}

func (*Car) TableName() string {
	return "cars"
}

func (l Car) Validate() error {
	v := validator.New()
	return v.Struct(l)
}
