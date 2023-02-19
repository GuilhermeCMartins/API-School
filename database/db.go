package database

import (
	"api-school/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDB() {
	stringConnection := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringConnection))
	if err != nil {
		log.Panic("Failed to connect to db.")
	}

	DB.AutoMigrate(&models.Teacher{}, &models.Student{}, &models.Classroom{}, &models.Grade{})
}
