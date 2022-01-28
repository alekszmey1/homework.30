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

//метод добавления друзей
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

func (r *repository) DeleteUser(user *entity.DeleteUser) string {
	b, err := strconv.Atoi(user.TargetId)
	if err != nil {
		fmt.Println("err")
	}
	var name string
	for id, value := range r.usersById {
		if id == b {
			c := r.usersById[id]
			name = c.Name
		}
		for id2, value2 := range value.Friends {
			if value2 == b {
				value.Friends = append(value.Friends[:id2], value.Friends[id2+1:]...)
			}
		}
	}
	delete(r.usersById, b)
	return name
}
