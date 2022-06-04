package notification

import "gitlab.com/sholludev/sampoerna_notification/app/kategori"

type SingleNotifRequestDTO struct {
	FCMToken string      `json:"fcm_token" validate:"required"`
	Type     string      `json:"type" validate:"required"`
	Message  string      `json:"message"`
	Data     interface{} `json:"data"`
}

type NotifikasiResponseDTO struct {
	ID         uint                          `json:"id"`
	IDKategori uint                          `json:"id_kategori"`
	Kategori   *kategori.KategoriResponseDTO `json:"kategori"`
	IDUser     uint                          `json:"id_user"`
	Judul      string                        `json:"judul"`
	Deskripsi  string                        `json:"deskripsi"`
	IsRead     bool                          `json:"is_read"`
	IsActive   bool                          `json:"is_active"`
	CreatedAt  string                        `json:"created_at"`
}
