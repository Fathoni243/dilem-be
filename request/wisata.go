package request

type (
	WisataRequest struct {
		Nama       string `form:"nama" binding:"required"`
		HargaTiket int    `form:"harga_tiket" binding:"required"`
		Alamat     string `form:"alamat" binding:"required"`
		Kontak     string `form:"kontak" binding:"required"`
		Gambar     string 
		JamBuka    string `form:"jam_buka" binding:"required"`
		JamTutup   string `form:"jam_tutup" binding:"required"`
		Deskripsi  string `form:"deskripsi" binding:"required"`
	}
)
