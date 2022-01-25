package repository

import (
	"fmt"
	"homework.30/pkg/entity"
	"strconv"
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

func (r *repository) MakeFriends(friends *entity.MakeFriends) (a int, b int, err error) {
	a, err = makeFriends(friends.SourceId)
	b, err = makeFriends(friends.TargetId)
	user1 := r.usersById[a]
	user2 := r.usersById[b]
	user1.Friends = append(user1.Friends, b)
	user2.Friends = append(user2.Friends, a)
	return a, b, err
}

func makeFriends(s string) (int, error) {
	a, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("err")
	}
	return a, err
}
