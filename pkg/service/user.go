package service

import "sayakaya-test/pkg/entity"

type UserService struct {
	UserRepository entity.UserRepository
}

func NewUserService(ur entity.UserRepository) entity.UserService {
	return &UserService{
		UserRepository: ur,
	}
}

func (us *UserService) FetchUsers() ([]entity.User, error) {
	return us.UserRepository.FetchUsers()
}

func (us *UserService) UpdateIsBirthday(id int, isBirthday bool) error {
	return us.UserRepository.UpdateIsBirthday(id, isBirthday)
}
