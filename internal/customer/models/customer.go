package models

import (
	"gorm.io/gorm"
	"time"
)

type Customer struct {
	Id        string `gorm:"primarykey"`
	Name      string
	Phone     string
	Address   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type RequestAddCustomer struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}
