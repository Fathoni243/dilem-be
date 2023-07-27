package service

import (
	"web-desa/model"
	"web-desa/request"

	supabasestorageuploader "github.com/adityarizkyramadhan/supabase-storage-uploader"
	"github.com/gin-gonic/gin"
)

type wisataService struct {
	wisataRepository model.WisataRepository
}

func NewWisataService(wisata model.WisataRepository) model.WisataService {
	return &wisataService{wisataRepository: wisata}
}

// Destroywisata implements model.WisataService
func (w *wisataService) DestroyWisata(id uint) error {
	wisata, _ := w.wisataRepository.FindByID(id)

	err := w.wisataRepository.Delete(wisata)
	if err != nil {
		return err
	}
	return nil
}

// Editwisata implements model.WisataService
func (w *wisataService) EditWisata(id uint, req *request.WisataRequest) (*model.Wisata, error) {
	newWisata, err := w.wisataRepository.UpdateByID(&model.Wisata{
		ID:         id,
		Nama:       req.Nama,
		HargaTiket: req.HargaTiket,
		Alamat:     req.Alamat,
		Kontak:     req.Kontak,
		Gambar:     req.Gambar,
		JamBuka:    req.JamBuka,
		JamTutup:   req.JamTutup,
		Deskripsi:  req.Deskripsi,
	})

	if err != nil {
		return nil, err
	}

	return newWisata, err
}

// Fetchwisata implements model.WisataService
func (w *wisataService) FetchWisata() ([]*model.Wisata, error) {
	wisatas, err := w.wisataRepository.Fetch()
	if err != nil {
		return nil, err
	}

	return wisatas, err
}

// GetByID implements model.WisataService
func (w *wisataService) GetByID(id uint) (*model.Wisata, error) {
	matkul, err := w.wisataRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	return matkul, err
}

// Storewisata implements model.WisataService
func (w *wisataService) StoreWisata(req *request.WisataRequest) (*model.Wisata, error) {
	wisata := &model.Wisata{
		Nama:       req.Nama,
		HargaTiket: req.HargaTiket,
		Alamat:     req.Alamat,
		Kontak:     req.Kontak,
		Gambar:     req.Gambar,
		JamBuka:    req.JamBuka,
		JamTutup:   req.JamTutup,
		Deskripsi:  req.Deskripsi,
	}

	newwisata, err := w.wisataRepository.Create(wisata)
	if err != nil {
		return nil, err
	}

	return newwisata, nil
}

var supClientWisata = supabasestorageuploader.NewSupabaseClient(
	"https://qbekdjxviuehumdvbstm.supabase.co",
	"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6InFiZWtkanh2aXVlaHVtZHZic3RtIiwicm9sZSI6ImFub24iLCJpYXQiOjE2OTAyMDI4MTYsImV4cCI6MjAwNTc3ODgxNn0.BNktDUzBhCu8bfALLtQ7LxcZnrWVeJwXjh8S3I_iP0E",
	"api-service-desa",
	"wisata",
)

func (w *wisataService) UploadImage(c *gin.Context) (string, error) {
	file, err := c.FormFile("gambar")
	if err != nil {
		return "", err
	}
	link, err := supClientWisata.Upload(file)
	if err != nil {
		return "", err
	}
	return link, nil
}

func (w *wisataService) DeleteImage(c *gin.Context, id uint) error {
	infoKegiatan, errFind := w.GetByID(id)
	if errFind != nil {
		return errFind
	}

	_, err := supClientWisata.DeleteFile(infoKegiatan.Gambar)
	if err != nil {
		return err
	} 

	return nil
}