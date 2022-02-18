package db

import (
	"Jokaru-py/managingEntities/models"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	serverAdress = "0.0.0.0:8100"
	HOST         = "localhost"
	TimeZone     = "Europe/Moscow"
	PORT         = "12016"
)

func New() *gorm.DB {
	dbUser, dbPassword, dbName, HOST, PORT :=
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("HOST_DB"),
		os.Getenv("PORT_DB")
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		HOST, PORT, dbUser, dbPassword, dbName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("storage err: ", err)
	}

	return db
}

//TODO: проверка на ошибки
func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&models.UsersDB{},  // Таблицы с данными пользователей
		&models.ObjectDB{}, // Таблица с данными объектов
	)
}
