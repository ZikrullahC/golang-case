package main

import (
	"vatan-soft-go-staj-case/db"
	"vatan-soft-go-staj-case/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Veritabanı bağlantısını başlatın
	db.Connect()

	// Echo instance'ı oluşturun
	e := echo.New()

	// Logger Middleware:
	e.Use(middleware.Logger())

	// Recover Middleware:
	e.Use(middleware.Recover())

	// Global veritabanı bağlantısını middleware ile ekle
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("db", db.DB)
			return next(c)
		}
	})

	// Rotaları tanımlayın
	e.POST("/ogrenci", handler.OgrenciOlustur)
	e.GET("/ogrenci/:id", handler.OgrenciBilgisiAl)
	e.PUT("/ogrenci/:id", handler.OgrenciGuncelle)
	e.DELETE("/ogrenci/:id", handler.OgrenciSil)
	e.GET("/ogrenci", handler.TumOgrenciBilgileriAl)

	e.POST("/plan", handler.PlanOlustur)
	e.GET("/plan/:id", handler.PlanBilgisiAl)
	e.PUT("/plan/:id", handler.PlanGuncelle)
	e.DELETE("/plan/:id", handler.PlanSil)
	e.GET("/plan/ogrenci/:id", handler.AylikPlanlar) // baslangic_gunu query parametresi olarak gönderilmeli

	// Sunucuyu başlatın
	e.Logger.Fatal(e.Start(":8080"))
}
