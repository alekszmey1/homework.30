package usecase

import (
	"homework.30/pkg/entity"
)

type (
	Usecase interface {
		CreateUser(*entity.User) (int, error)
		// DeleteUser(int) error
		// UpdateUser(int, int) error
		// GetFriends(int) ([]int, error)
		// MakeFriends(int, int)
	}

	Repository interface {
		CreateUser(*entity.User) (int, error)
		// DeleteUser(int) error
		// UpdateAge(int, int) error
		// MakeFriends(int, int) (string, string, error)
		// GetFriends(int) ([]string, error)
	}
)

type usecase struct {
	repository Repository
}

func NewUsecase(repository Repository) *usecase {
	return &usecase{
		repository: repository,
	}
}

// CreateUser return entity id or error
func (u *usecase) CreateUser(user *entity.User) (int, error) {
	uid, error := u.repository.CreateUser(user)
	return uid, error
}