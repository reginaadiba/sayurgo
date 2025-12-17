package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "host=localhost user=postgres password=123456 dbname=sayurgo port=5432 sslmode=disable TimeZone=Asia/Jakarta"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Gagal connect DB:", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("❌ Gagal ambil sqlDB:", err)
	}

	if err := sqlDB.Ping(); err != nil {
		log.Fatal("❌ Ping DB gagal:", err)
	}

	DB = db
	log.Println("✅ Database connected")
}
