package usecase

import (
	"homework.30/pkg/entity"
)

type (
	Usecase interface {
		CreateUser(*entity.User) (int, error)
		MakeFriends(*entity.MakeFriends) (int, int, error)
		DeleteUser(user *entity.DeleteUser) string
		// UpdateUser(int, int) error
		// GetFriends(int) ([]int, error)
	}

	Repository interface {
		CreateUser(*entity.User) (int, error)
		DeleteUser(user *entity.DeleteUser) string
		// UpdateAge(int, int) error
		MakeFriends(*entity.MakeFriends) (int, int, error)
		//GetFriends(int) ([]string, error)
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

func (u *usecase) MakeFriends(friends *entity.MakeFriends) (a, b int, err error) {
	a, b, err = u.repository.MakeFriends(friends)
	return a, b, err
}

func (u *usecase) DeleteUser(user *entity.DeleteUser) string {
	b := u.repository.DeleteUser(user)
	return b
}
