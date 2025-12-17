package models

import "time"

type Product struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"type:varchar(100);not null"`
	Price     int    `gorm:"not null"`
	Stock     int    `gorm:"not null"`
	Image     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
