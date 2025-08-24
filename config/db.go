package config

import (
	"fmt"
	"go-initial-project/entity"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		AppConfig.DB.Host,
		AppConfig.DB.User,
		AppConfig.DB.Pass,
		AppConfig.DB.Name,
		AppConfig.DB.Port,
		AppConfig.DB.SSLMode,
		AppConfig.DB.TimeZone,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database:", err)
	}

	sqlDB, _ := db.DB()
	sqlDB.Exec("SET TIME ZONE ?", AppConfig.DB.TimeZone)

	err = db.AutoMigrate(&entity.User{}, &entity.Activity{})
	if err != nil {
		return nil
	}

	return db
}
