package model

import (
	"time"
	"web-desa/request"

	"github.com/gin-gonic/gin"
)

type (
	Wisata struct {
		ID         uint      `json:"id" gorm:"primaryKey"`
		Nama       string    `json:"nama" gorm:"type:varchar(200)"`
		HargaTiket int       `json:"harga_tiket"`
		Alamat     string    `json:"alamat" gorm:"type:varchar(200)"`
		Kontak     string    `json:"kontak" gorm:"type:varchar(15)"`
		Gambar     string    `json:"gambar"`
		JamBuka    string    `json:"jam_buka"`
		JamTutup   string    `json:"jam_tutup"`
		Deskripsi  string    `json:"deskripsi"`
		CreatedAt  time.Time `json:"-"`
		UpdatedAt  time.Time `json:"-"`
	}

	WisataRepository interface {
		Create(wisata *Wisata) (*Wisata, error)
		UpdateByID(wisata *Wisata) (*Wisata, error)
		FindByID(id uint) (*Wisata, error)
		Delete(wisata *Wisata) error
		Fetch() ([]*Wisata, error)
	}

	WisataService interface {
		StoreWisata(req *request.WisataRequest) (*Wisata, error)
		EditWisata(id uint, req *request.WisataRequest) (*Wisata, error)
		GetByID(id uint) (*Wisata, error)
		DestroyWisata(id uint) error
		FetchWisata() ([]*Wisata, error)
		UploadImage(c *gin.Context) (string, error)
		DeleteImage(c *gin.Context, id uint) error
	}

	WisataHandler interface {
		Mount(group *gin.RouterGroup)
	}
)
