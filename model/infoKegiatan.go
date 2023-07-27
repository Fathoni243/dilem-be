package model

import (
	"time"
	"web-desa/request"

	"github.com/gin-gonic/gin"
)

type (
	InfoKegiatan struct {
		ID        	uint   	`gorm:"primaryKey"`
		Judul     	string 	`json:"tentang_desa" gorm:"type:varchar(100)"`
		Gambar    	string 	`json:"gambar" gorm:"type:varchar(200)"`
		Tanggal   	string 	`json:"tanggal" gorm:"type:varchar(20)"`
		Deskripsi 	string 	`json:"deskripsi"`
		CreatedAt 	time.Time `json:"-"`
		UpdatedAt 	time.Time `json:"-"`
	}

	InfoKegiatanRepository interface {
		Create(infoKegiatan *InfoKegiatan) (*InfoKegiatan, error)
		UpdateByID(infoKegiatan *InfoKegiatan) (*InfoKegiatan, error)
		FindByID(id uint) (*InfoKegiatan, error)
		Delete(infoKegiatan *InfoKegiatan) (*InfoKegiatan, error)
		Fetch() ([]*InfoKegiatan, error)
	}

	InfoKegiatanService interface {
		StoreInfoKegiatan(req *request.InfoKegiatanRequest) (*InfoKegiatan, error)
		EditInfoKegiatan(id uint, req *request.InfoKegiatanRequest) (*InfoKegiatan, error)
		GetByID(id uint) (*InfoKegiatan, error)
		DestroyInfoKegiatan(id uint) error
		FetchInfoKegiatan() ([]*InfoKegiatan, error)
		UploadImage(c *gin.Context) (string, error)
		DeleteImage(c *gin.Context, id uint) error
	}

	InfoKegiatanHandler interface {
		Mount(group *gin.RouterGroup)
	}
)