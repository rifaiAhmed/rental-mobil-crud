package models

import "time"

type UriId struct {
	ID int `uri:"id" binding:"required"`
}

type OrderParam struct {
	ID              uint      `json:"id"`
	CarID           uint      `json:"car_id"`
	OrderDate       time.Time `json:"order_date"`
	PickupDate      time.Time `json:"pickup_date"`
	PickupLocation  string    `json:"pickup_location"`
	DropoffLocation string    `json:"dropoff_location"`

	Car Car `gorm:"foreignKey:CarID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type ComponentServerSide struct {
	Limit     int    `json:"limit"`
	Skip      int    `json:"skip"`
	SortType  string `json:"sort_type"`
	SortBy    string `json:"sort_by"`
	Search    string `json:"search"`
	Offset    int    `json:"offset"`
	Condition string `json:"condition"`
	From      string `json:"from"`
	To        string `json:"to"`
}
