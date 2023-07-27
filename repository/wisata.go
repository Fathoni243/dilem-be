package repository

import (
	"web-desa/config"
	"web-desa/model"
)

type wisataRepository struct {
	cfg config.Config
}

func NewWisataRepository(cfg config.Config) model.WisataRepository {
	return &wisataRepository{cfg: cfg}
}

func (w *wisataRepository) Create(wisata *model.Wisata) (*model.Wisata, error) {
	err := w.cfg.Database().Create(&wisata).Error
	if err != nil {
		return nil, err
	}

	return wisata, err
}

// Delete implements model.wisataRepository
func (w *wisataRepository) Delete(wisata *model.Wisata) error {
	err := w.cfg.Database().Delete(&wisata).Error
	if err != nil {
		return err
	}

	return nil
}

// Fetch implements model.wisataRepository
func (w *wisataRepository) Fetch() ([]*model.Wisata, error) {
	var data []*model.Wisata

	err := w.cfg.Database().Find(&data).Error
	if err != nil {
		return nil, err
	}

	return data, err
}

// FindByID implements model.WisataRepository
func (w *wisataRepository) FindByID(id uint) (*model.Wisata, error) {
	wisata := new(model.Wisata)

	err := w.cfg.Database().First(wisata, id).Error
	if err != nil {
		return nil, err
	}

	return wisata, err
}

// UpdateByID implements model.WisataRepository
func (w *wisataRepository) UpdateByID(wisata *model.Wisata) (*model.Wisata, error) {
	err := w.cfg.Database().Model(&wisata).Updates(&wisata).First(&wisata).Error
	if err != nil {
		return nil, err
	}
	
	return wisata, err
}