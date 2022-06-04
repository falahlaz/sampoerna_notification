package kategori

type KategoriResponseDTO struct {
	ID        uint   `json:"id"`
	Kode      string `json:"kode"`
	Nama      string `json:"nama"`
	IsActive  bool   `json:"is_active"`
	CreatedAt string `json:"created_at"`
}
