package service

import (
	"web-desa/handler"
	"web-desa/model"
	"web-desa/request"
)

type userService struct {
	userRepository model.UserRepository
}

func NewUserService(repo model.UserRepository) *userService {
	return &userService{userRepository: repo}
}

func (s *userService) Register(userRequest *request.UserRequest) (*model.User, error) {
	user := &model.User{
		Username: userRequest.Username,
		Password: handler.HashPassword(userRequest.Password),
	}

	newUser, err := s.userRepository.Create(user)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}

func (s *userService) GetAllUser() ([]*model.User, error) {
	users, err := s.userRepository.ReadAll()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *userService) GetByUsername(Username string) (*model.User, error) {
	user, err := s.userRepository.FindByUsername(Username)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) EditUser(id uint, user *request.UserRequest) (*model.User, error) {
	newUser, err := s.userRepository.Update(&model.User{
		ID: id,
		Username: user.Username,
		Password: handler.HashPassword(user.Password),
	})

	if err != nil {
		return nil, err
	}

	return newUser, nil
}

func (s *userService) DestroyUser(id uint) error {
	user, errFind := s.userRepository.FindById(id)
	if errFind != nil {
		return errFind
	}

	err := s.userRepository.Delete(user)
	if err != nil {
		return err
	}

	return nil
}