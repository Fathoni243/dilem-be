package repository

import (
	"web-desa/config"
	"web-desa/model"
)

type userRepository struct {
	cfg config.Config
}

func NewUserRepository(cfg config.Config) model.UserRepository{
	return &userRepository{cfg: cfg}
}

func (u *userRepository) Create(user *model.User) (*model.User, error){
	err:= u.cfg.Database().Create(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userRepository) FindById(id uint) (*model.User, error) {
	user := new(model.User)

	err := u.cfg.Database().First(user, id).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userRepository) FindByUsername(username string) (*model.User, error) {
	var user *model.User

	err := u.cfg.Database().First(&user, "username = ?", username).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userRepository) ReadAll() ([]*model.User, error){
	var users []*model.User

	err := u.cfg.Database().Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u *userRepository) Update(user *model.User) (*model.User, error){
	err := u.cfg.Database().Model(&user).Updates(&user).First(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userRepository) Delete(user *model.User) error  {
	err := u.cfg.Database().Delete(&user).Error
	if err != nil {
		return err
	}

	return nil
}