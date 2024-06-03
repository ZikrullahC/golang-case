package model

import "gorm.io/gorm"

type Ogrenci struct {
	gorm.Model

	Isim    string `json:"isim"`
	Email   string `json:"email"`
	Sifre   string `json:"sifre"`
	Planlar []Plan `json:"etkinlikler" gorm:"foreignKey:OgrenciID"`
}
