package app

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"model"
)

func dbConnection() *gorm.DB {
	dbUser := "ewlight"
	dbPass := "mysoul3nity"
	dbHost := "localhost"
	dbName := "xiomicamp"
	dbPort := "5432"

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", dbHost, dbUser, dbPass, dbName, dbPort)

	db, errorDB := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if errorDB != nil {
		panic("Error Connecting Database")
	}
	return db
}

func dbMigration(db *gorm.DB) {
	db.AutoMigrate(&model.Order{}, &model.Item{})
	fmt.Println("DB Migration Done")
}

func dbDisconnect(db *gorm.DB)  {
	sqlDB, err := db.DB()
	if err != nill { panic("dbConnection Not Found")}
	sqlDB.Close()
}

