package handler

import (
	"net/http"
	"vatan-soft-go-staj-case/db"
	"vatan-soft-go-staj-case/model"

	"github.com/labstack/echo/v4"
)

func OgrenciOlustur(c echo.Context) error {
	ogrenci := new(model.Ogrenci)

	if err := c.Bind(ogrenci); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	if result := db.DB.Create(&ogrenci); result.Error != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": result.Error.Error()})
	}
	return c.JSON(http.StatusCreated, ogrenci)
}

func OgrenciBilgisiAl(c echo.Context) error {
	id := c.Param("id")
	var ogrenci model.Ogrenci

	if err := db.DB.First(&ogrenci, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Ogrenci bulunamadi"})
	}

	return c.JSON(http.StatusOK, ogrenci)
}

func OgrenciGuncelle(c echo.Context) error {
	id := c.Param("id")
	var ogrenci model.Ogrenci

	// First metotu ogrenci degiskenine id degerindeki ogrencinin bilgilerini doldurur
	if err := db.DB.First(&ogrenci, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Ogrenci bulunamadi"})
	}
	if err := c.Bind(&ogrenci); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	// Save fonksiyonu sadece mevcut kaydi guncellemek icin kullanilir.
	if sonuc := db.DB.Save(&ogrenci); sonuc.Error != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": sonuc.Error.Error()})
	}
	return c.JSON(http.StatusOK, ogrenci)
}

func OgrenciSil(c echo.Context) error {
	id := c.Param("id")

	if err := db.DB.Delete(&model.Ogrenci{}, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Ogrenci bulunamadi"})
	}

	return c.JSON(http.StatusNoContent, echo.Map{"mesaj": "Ogrenci basarili bir sekilde silindi"})
}

func TumOgrenciBilgileriAl(c echo.Context) error {
	var ogrenciler []model.Ogrenci

	if err := db.DB.Find(&ogrenciler).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, ogrenciler)
}
