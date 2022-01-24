package repository

import (
	"homework.30/pkg/entity"
)

//Определение хранилища, и типы данных, хранящиеся там
type repository struct {
	index     int
	usersById map[int]*entity.User
}

// Функция создания хранилища
func NewRepository() *repository {
	return &repository{
		usersById: make(map[int]*entity.User),
	}
}

// Метод добавления в хранилище новую сущность
func (r *repository) CreateUser(user *entity.User) (int, error) {
	r.index++
	user.Id = r.index
	r.usersById[user.Id] = user
	return user.Id, nil
}
