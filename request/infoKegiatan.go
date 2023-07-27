package request

type (
	InfoKegiatanRequest struct {
		Judul     string `form:"judul" binding:"required"`
		Gambar    string 
		Tanggal   string `form:"tanggal" binding:"required"`
		Deskripsi string `form:"deskripsi" binding:"required"`
	}
)