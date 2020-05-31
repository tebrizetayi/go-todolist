package models

import (
	"time"
)

// MyGormModel mimixks GormModel but uses uuid's for ID, generated in go
type MyGormModel struct {
	ID        int `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}