package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Email      string    `json:"email" gorm:"unique;not null"`
	Password   string    `json:"password" gorm:"not null"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	Active     bool      `json:"active" gorm:"default:true"`
	DateJoined time.Time `json:"date_joined" gorm:"autoCreateTime"`
}
