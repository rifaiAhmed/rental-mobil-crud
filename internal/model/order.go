package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Order struct {
	ID              uint      `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	CarID           uint      `json:"car_id" gorm:"not null"`
	OrderDate       time.Time `json:"order_date" gorm:"type:date;not null"`
	PickupDate      time.Time `json:"pickup_date" gorm:"type:date;not null"`
	PickupLocation  string    `json:"pickup_location" gorm:"type:varchar(255);not null"`
	DropoffLocation string    `json:"dropoff_location" gorm:"type:varchar(255);not null"`

	Car Car `gorm:"foreignKey:CarID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (*Order) TableName() string {
	return "orders"
}

func (l Order) Validate() error {
	v := validator.New()
	return v.Struct(l)
}
