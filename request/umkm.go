package request

type (
	UmkmRequest struct {
		Nama       string `form:"nama" binding:"required"`
		Alamat     string `form:"alamat" binding:"required"`
		Kontak     string `form:"kontak" binding:"required"`
		Gambar     string 
		Deskripsi  string `form:"deskripsi" binding:"required"`
	}
)
