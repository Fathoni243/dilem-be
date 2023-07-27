package service

import (
	"web-desa/model"
	"web-desa/request"

	supabasestorageuploader "github.com/adityarizkyramadhan/supabase-storage-uploader"
	"github.com/gin-gonic/gin"
)

type umkmService struct {
	umkmRepository model.UmkmRepository
}


func NewumkmService(umkm model.UmkmRepository) model.UmkmService {
	return &umkmService{umkmRepository: umkm}
}

// DestroyUmkm implements model.UmkmService
func (u *umkmService) DestroyUmkm(id uint) error {
	umkm, _ := u.umkmRepository.FindByID(id)
	
	err := u.umkmRepository.Delete(umkm)
	if err != nil {
		return err
	}
	return nil
}

// EditUmkm implements model.UmkmService
func (u *umkmService) EditUmkm(id uint, req *request.UmkmRequest) (*model.Umkm, error) {
	newUmkm, err := u.umkmRepository.UpdateByID(&model.Umkm{
		ID:        id,
		Nama:      req.Nama,
		Alamat:    req.Alamat,
		Kontak:    req.Kontak,
		Gambar:    req.Gambar,
		Deskripsi: req.Deskripsi,
	})

	if err != nil {
		return nil, err
	}

	return newUmkm, err
}

// FetchUmkm implements model.UmkmService
func (u *umkmService) FetchUmkm() ([]*model.Umkm, error) {
	umkms, err := u.umkmRepository.Fetch()
	if err != nil {
		return nil, err
	}

	return umkms, err
}

// GetByID implements model.UmkmService
func (u *umkmService) GetByID(id uint) (*model.Umkm, error) {
	matkul, err := u.umkmRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	return matkul, err
}

// StoreUmkm implements model.UmkmService
func (u *umkmService) StoreUmkm(req *request.UmkmRequest) (*model.Umkm, error) {
	umkm := &model.Umkm{
		Nama:      req.Nama,
		Alamat:    req.Alamat,
		Kontak:    req.Kontak,
		Gambar:    req.Gambar,
		Deskripsi: req.Deskripsi,
	}

	newUmkm, err := u.umkmRepository.Create(umkm)
	if err != nil {
		return nil, err
	}

	return newUmkm, nil
}

var supClientUmkm = supabasestorageuploader.NewSupabaseClient(
	"https://qbekdjxviuehumdvbstm.supabase.co",
	"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6InFiZWtkanh2aXVlaHVtZHZic3RtIiwicm9sZSI6ImFub24iLCJpYXQiOjE2OTAyMDI4MTYsImV4cCI6MjAwNTc3ODgxNn0.BNktDUzBhCu8bfALLtQ7LxcZnrWVeJwXjh8S3I_iP0E",
	"api-service-desa",
	"umkm",
)

func (h *umkmService) UploadImage(c *gin.Context) (string, error) {
	file, err := c.FormFile("gambar")
	if err != nil {
		return "", err
	}
	link, err := supClientUmkm.Upload(file)
	if err != nil {
		return "", err
	}
	return link, nil
}

func (h *umkmService) DeleteImage(c *gin.Context, id uint) error {
	umkm, errFind := h.GetByID(id)
	if errFind != nil {
		return errFind
	}

	_, err := supClientUmkm.DeleteFile(umkm.Gambar)
	if err != nil {
		return err
	} 

	return nil
}