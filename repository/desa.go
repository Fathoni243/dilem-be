package repository

import (
	"web-desa/config"
	"web-desa/model"
)

type desaRepository struct {
	cfg config.Config
}

func NewDesaRepository(cfg config.Config) model.DesaRepository {
	return &desaRepository{cfg: cfg}
}

func (r *desaRepository) Create(desa *model.Desa) (*model.Desa, error) {
	err := r.cfg.Database().Create(&desa).Error
	if err != nil {
		return nil, err
	}

	return desa, nil
}

func (r *desaRepository) Fetch() (*model.Desa, error) {
	var data *model.Desa

	err := r.cfg.Database().First(&data).Error
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (r *desaRepository) Update(desa *model.Desa) (*model.Desa, error) {
	err := r.cfg.Database().Model(&desa).Updates(&desa).First(&desa).Error
	if err != nil {
		return nil, err
	}

	return desa, nil
}

func (r *desaRepository) Delete(desa *model.Desa) (*model.Desa, error) {
	err := r.cfg.Database().Delete(&desa).Error
	if err != nil {
		return nil, err
	}

	return desa, nil
}