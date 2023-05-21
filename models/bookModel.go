package models

import "gorm.io/gorm"

type Books struct {
	gorm.Model
	Judul         string
	Penulis       string
	Penerbit      string
	Tanggal_rilis string
}
