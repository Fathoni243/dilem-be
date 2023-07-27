package service

import (
	"web-desa/model"
	"web-desa/request"
)

type desaService struct {
	desaRepository model.DesaRepository
}

func NewDesaService(desa model.DesaRepository) model.DesaService {
	return &desaService{desaRepository: desa}
}

func (s *desaService) StoreDesa(req *request.DesaRequest) (*model.Desa, error) {
	desa := &model.Desa{
		ID: 1,
		TentangDesa: req.TentangDesa,
		Visi: req.Visi,
		Misi: req.Misi,
	}

	newDesa, err := s.desaRepository.Create(desa)
	if err != nil {
		return nil, err
	}

	return newDesa, nil
}

func (s *desaService) FetchDesa() (*model.Desa, error) {
	desa, err := s.desaRepository.Fetch()
	if err != nil {
		return nil, err
	}

	return desa, nil
}

func (s *desaService) EditDesa(id uint, req *request.DesaRequest) (*model.Desa, error) {
	newDesa, err := s.desaRepository.Update(&model.Desa{
		ID: id,
		TentangDesa: req.TentangDesa,
		Visi: req.Visi,
		Misi: req.Misi,
	})

	if err != nil {
		return nil, err
	}

	return newDesa, nil
}

func (s *desaService) DestroyDesa() error {
	desa, _ := s.desaRepository.Fetch()

	_, err := s.desaRepository.Delete(desa)
	if err != nil {
		return err
	}

	return nil
}