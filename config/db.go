package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func ConnectDB() {
	user := GetEnv("DB_USER", "root")
	pass := GetEnv("DB_PASS", " ")
	host := GetEnv("DB_HOST", "127.0.0.1")
	port := GetEnv("DB_PORT", "3306")
	name := GetEnv("DB_NAME", "obsidian")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, pass, host, port, name)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	DB = db
	log.Println("Connected to database")
}
