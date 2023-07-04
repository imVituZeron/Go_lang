package database

import (
	"log"
	"pack/api-gin/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDatabase() {
	Conn := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(Conn))
	if err != nil {
		log.Panic("Error try connect in database")
	}
	DB.AutoMigrate(&models.Aluno{})
}
