package repository

import (
	"example.com/golang-todo/internal/model"
)

type UserRepository struct {
	userModel *model.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		userModel: &model.User{},
	}
}

func (handler *UserRepository) GetAll() {
}

func (handler *UserRepository) GetOneById() {
}

func (handler *UserRepository) Create() {
}

func (handler *UserRepository) Update() {
}

func (handler *UserRepository) Delete() {
}
