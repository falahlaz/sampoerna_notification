package models

import "time"

type MKategori struct {
	ID         uint      `json:"id" gorm:"primary_key"`
	Kode       string    `json:"kode" gorm:"type:varchar(10);unique_index"`
	Nama       string    `json:"nama" gorm:"type:varchar(50)"`
	Keterangan string    `json:"keterangan" gorm:"type:varchar(100)"`
	IsActive   bool      `json:"is_active" gorm:"type:boolean;default:true"`
	CreatedBy  uint      `json:"created_by"`
	UpdatedBy  uint      `json:"updated_by"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
