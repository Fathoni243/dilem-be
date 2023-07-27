package repository

import (
	"web-desa/config"
	"web-desa/model"
)

type umkmRepository struct {
	cfg config.Config
}

func NewUmkmRepository(cfg config.Config) model.UmkmRepository {
	return &umkmRepository{cfg: cfg}
}

func (u *umkmRepository) Create(umkm *model.Umkm) (*model.Umkm, error) {
	err := u.cfg.Database().Create(&umkm).Error
	if err != nil {
		return nil, err
	}

	return umkm, err
}

// Delete implements model.UmkmRepository
func (u *umkmRepository) Delete(umkm *model.Umkm) error {
	err := u.cfg.Database().Delete(&umkm).Error
	if err != nil {
		return err
	}

	return nil
}

// Fetch implements model.UmkmRepository
func (u *umkmRepository) Fetch() ([]*model.Umkm, error) {
	var data []*model.Umkm

	err := u.cfg.Database().Find(&data).Error
	if err != nil {
		return nil, err
	}

	return data, err
}

// FindByID implements model.UmkmRepository
func (u *umkmRepository) FindByID(id uint) (*model.Umkm, error) {
	umkm := new(model.Umkm)

	err := u.cfg.Database().First(umkm, id).Error
	if err != nil {
		return nil, err
	}

	return umkm, err
}

// UpdateByID implements model.UmkmRepository
func (u *umkmRepository) UpdateByID(umkm *model.Umkm) (*model.Umkm, error) {
	err := u.cfg.Database().Model(&umkm).Updates(&umkm).First(&umkm).Error
	if err != nil {
		return nil, err
	}
	
	return umkm, err
}