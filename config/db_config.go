package config

import (
	"auth_microservice/entity"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var Db *gorm.DB

func DbConfig() {
	dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=%s",
		os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))
	var err error
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Couldn't connect !!!!")
	}
	err = Db.AutoMigrate(&entity.User{})
	err = Db.AutoMigrate(&entity.Token{})

	if err != nil {
		fmt.Println("Couldn't Migrate !!!!")
	}

}
