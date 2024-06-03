package model

import (
	"time"

	"gorm.io/gorm"
)

type Plan struct {
	// Bir structi veri tabani olarak olutururken gereklidir.
	gorm.Model

	PlanAdi         string    `json:"plan_adi"`
	Gun             string    `json:"gun"`
	BaslangicZamani time.Time `json:"baslangic_zamani"`
	BitisZamani     time.Time `json:"bitis_zamani"`
	Durum           string    `json:"durum"`
	OgrenciID       uint      `json:"ogrenci_id"`
}
