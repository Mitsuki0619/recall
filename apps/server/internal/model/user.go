package model

import "time"

type User struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	Email     string    `json:"email" gorm:"unique;not null;varchar(100)"`
	Password  string    `json:"password" gorm:"not null;varchar(50)"`
	FirstName string    `json:"first_name" gorm:"not null;varchar(50)"`
	LastName  string    `json:"last_name" gorm:"not null;varchar(50)"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
