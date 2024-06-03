package db

import (
	"log"

	"vatan-soft-go-staj-case/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dbBaglantisi := "username:password@tcp(127.0.0.1:3306)/vatan_soft_go_staj_case?charset=utf8mb4&parseTime=True&loc=Local"

	var err error

	DB, err = gorm.Open(mysql.Open(dbBaglantisi), &gorm.Config{})
	if err != nil {
		log.Fatalf("Veritabanına bağlanılamadı: %v", err)
	}

	if err := DB.AutoMigrate(&model.Ogrenci{}, &model.Plan{}); err != nil {
		log.Fatalf("AutoMigrate hatasi cikti: %v", err)
	}
}
