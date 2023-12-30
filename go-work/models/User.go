package models

import "time"

type User struct {
	ID        string `gorm:"primaryKey"`
	Username  string
	Email     string
	Password  string
	Photo     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
