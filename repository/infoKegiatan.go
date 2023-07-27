package repository

import (
	"web-desa/config"
	"web-desa/model"
)

type infoKegiatanRepository struct {
	cfg config.Config
}

func NewInfoKegiatanRepository(cfg config.Config) model.InfoKegiatanRepository {
	return &infoKegiatanRepository{cfg: cfg}
}

func (r *infoKegiatanRepository) Create(infoKegiatan *model.InfoKegiatan) (*model.InfoKegiatan, error) {
	err := r.cfg.Database().Create(&infoKegiatan).Error
	if err != nil {
		return nil, err
	}

	return infoKegiatan, nil
}

func (r *infoKegiatanRepository) Fetch() ([]*model.InfoKegiatan, error) {
	var data []*model.InfoKegiatan

	err := r.cfg.Database().Find(&data).Error
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (r *infoKegiatanRepository) FindByID(id uint) (*model.InfoKegiatan, error){
	infoKegiatan := new(model.InfoKegiatan)

	err := r.cfg.Database().First(infoKegiatan, id).Error
	if err != nil {
		return nil, err
	}

	return infoKegiatan, nil
}

func (r *infoKegiatanRepository) UpdateByID(infoKegiatan *model.InfoKegiatan) (*model.InfoKegiatan, error) {
	err := r.cfg.Database().Model(&infoKegiatan).Updates(&infoKegiatan).First(&infoKegiatan).Error
	if err != nil {
		return nil, err
	}

	return infoKegiatan, nil
}

func (r *infoKegiatanRepository) Delete(infoKegiatan *model.InfoKegiatan) (*model.InfoKegiatan, error) {
	err := r.cfg.Database().Delete(&infoKegiatan).Error
	if err != nil {
		return nil, err
	}

	return infoKegiatan, nil
}