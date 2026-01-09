package service

import (
	"example.com/golang-todo/internal/repository"
)

type UserService struct {
	userRepository *repository.UserRepository
}

func NewUserService(repository *repository.UserRepository) UserServiceInterface {
	return &UserService{
		userRepository: repository,
	}
}

func (service *UserService) GetAll() {
	service.userRepository.GetAll()
}

func (service *UserService) GetOneById() {
}

func (service *UserService) Create() {
}

func (service *UserService) Update() {
}

func (service *UserService) Delete() {
}
