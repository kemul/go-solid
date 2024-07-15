package usecase

import (
	"go-solid/domain"
)

type userUsecase struct {
	userRepo domain.UserRepository
}

func NewUserUsecase(repo domain.UserRepository) domain.UserUsecase {
	return &userUsecase{userRepo: repo}
}

func (u *userUsecase) GetUserByID(id int64) (*domain.User, error) {
	return u.userRepo.GetByID(id)
}

func (u *userUsecase) CreateUser(user *domain.User) error {
	return u.userRepo.Store(user)
}
