package model

import (
	"time"
	"web-desa/request"

	"github.com/gin-gonic/gin"
)

type (
	Umkm struct {
		ID        uint      `json:"id" gorm:"primaryKey"`
		Nama      string    `json:"nama" gorm:"type:varchar(200)"`
		Alamat    string    `json:"alamat" gorm:"type:varchar(200)"`
		Kontak    string    `json:"kontak" gorm:"type:varchar(15)"`
		Gambar    string    `json:"gambar"`
		Deskripsi string    `json:"deskripsi"`
		CreatedAt time.Time `json:"-"`
		UpdatedAt time.Time `json:"-"`
	}

	UmkmRepository interface {
		Create(Umkm *Umkm) (*Umkm, error)
		UpdateByID(Umkm *Umkm) (*Umkm, error)
		FindByID(id uint) (*Umkm, error)
		Delete(Umkm *Umkm) error
		Fetch() ([]*Umkm, error)
	}

	UmkmService interface {
		StoreUmkm(req *request.UmkmRequest) (*Umkm, error)
		EditUmkm(id uint, req *request.UmkmRequest) (*Umkm, error)
		GetByID(id uint) (*Umkm, error)
		DestroyUmkm(id uint) error
		FetchUmkm() ([]*Umkm, error)
		UploadImage(c *gin.Context) (string, error)
		DeleteImage(c *gin.Context, id uint) error
	}

	UmkmHandler interface {
		Mount(group *gin.RouterGroup)
	}
)
