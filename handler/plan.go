package handler

import (
	"net/http"
	"time"
	"vatan-soft-go-staj-case/db"
	"vatan-soft-go-staj-case/model"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func PlanOlustur(c echo.Context) error {
	plan := new(model.Plan)

	if err := c.Bind(plan); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	db := c.Get("db").(*gorm.DB)

	var cakisanPlanlar []model.Plan
	if err := db.Where("ogrenci_id = ? AND ((baslangic_zamani <= ? AND bitis_zamani >= ?) OR (baslangic_zamani >= ? AND bitis_zamani <= ?))",
		plan.OgrenciID, plan.BaslangicZamani, plan.BitisZamani, plan.BaslangicZamani, plan.BitisZamani).Find(&cakisanPlanlar).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	mesaj := "Plan oluşturuldu"
	if len(cakisanPlanlar) > 0 {
		mesaj = "Plan olusturuldu ama cakisan planlar var"
	}

	if err := db.Create(&plan).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	cevap := map[string]interface{}{
		"mesaj": mesaj,
		"plan":  plan,
	}

	return c.JSON(http.StatusCreated, cevap)
}

func PlanBilgisiAl(c echo.Context) error {
	id := c.Param("id")
	var plan model.Plan

	if err := db.DB.First(&plan, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Plan bulunamadi"})
	}

	return c.JSON(http.StatusOK, plan)
}

func PlanGuncelle(c echo.Context) error {
	id := c.Param("id")
	var plan model.Plan

	if err := db.DB.First(&plan, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Plan bulunamadi"})
	}

	if err := c.Bind(&plan); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	if err := db.DB.Save(&plan).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, plan)
}

func PlanSil(c echo.Context) error {
	id := c.Param("id")

	if err := db.DB.Delete(&model.Plan{}, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Plan bulunamadi"})
	}

	return c.NoContent(http.StatusNoContent)
}

func PlanListele(c echo.Context) error {
	ogrenciID := c.Param("id")
	var planlar []model.Plan

	if err := db.DB.Where("ogrenci_id = ?", ogrenciID).Find(&planlar).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, planlar)
}

func HaftalikPlanlar(c echo.Context) error {
	ogrenciID := c.Param("id")
	var planlar []model.Plan

	baslangicGunuStr := c.QueryParam("baslangic_gunu")
	if baslangicGunuStr == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Baslangic gunu gerekli"})
	}

	dateFormat := "02-01-2006"
	baslangicGunu, err := time.Parse(dateFormat, baslangicGunuStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Gecersiz tarih formati. Dogru format: dd-MM-yyyy"})
	}

	bitisGunu := baslangicGunu.AddDate(0, 0, 7)

	err = db.DB.Where("ogrenci_id = ? AND baslangic_gunu >= ? AND bitis_gunu < ?", ogrenciID, baslangicGunu, bitisGunu).Find(&planlar).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, planlar)
}

func AylikPlanlar(c echo.Context) error {
	ogrenciID := c.Param("id")

	baslangicGunuStr := c.QueryParam("baslangic_gunu")
	if baslangicGunuStr == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Baslangic gunu gerekli"})
	}

	dateFormat := "02-01-2006"
	baslangicGunu, err := time.Parse(dateFormat, baslangicGunuStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Gecersiz tarih formati. Dogru format: dd-MM-yyyy"})
	}

	var planlar []model.Plan
	bitisGunu := baslangicGunu.AddDate(0, 0, 30)

	err = db.DB.Where("ogrenci_id = ? AND baslangic_gunu >= ? AND bitis_gunu < ?", ogrenciID, baslangicGunu, bitisGunu).Find(&planlar).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, planlar)
}
