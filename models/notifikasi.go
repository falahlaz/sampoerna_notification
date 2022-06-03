package models

import "time"

type TNotifikasi struct {
	ID            uint      `json:"id" gorm:"primary_key"`
	IDKategori    uint      `json:"id_kategori"`
	IDUser        uint      `json:"id_user"`
	OriginService string    `json:"origin_service" gorm:"not null"`
	Judul         string    `json:"judul"`
	Deskripsi     string    `json:"deskripsi"`
	IsRead        bool      `json:"is_read" gorm:"default:false"`
	IsActive      bool      `json:"is_active" gorm:"default:true"`
	CreatedBy     uint      `json:"created_by"`
	UpdatedBy     uint      `json:"updated_by"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
