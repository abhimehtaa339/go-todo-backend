package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Username   string    `json:"username" gorm:"unique;not null"`
	Password   string    `json:"password" gorm:"not null"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	Email      string    `json:"email" gorm:"unique;not null"`
	Active     bool      `json:"active" gorm:"default:true"`
	DateJoined time.Time `json:"date_joined" gorm:"autoCreateTime"`
}
