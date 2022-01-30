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
	a := stringToInt(s)
	return a, nil
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

func (r *repository) GetFriends(a int) (b []string, err error) {
	for id, _ := range r.usersById {
		if id == a {
			c := r.usersById[id]
			d := c.Friends
			for _, idFriends := range d {
				for id, user := range r.usersById {
					if idFriends == id {
						b = append(b, user.Name)
					}
				}
			}
		}
	}
	return b, nil
}

func (r *repository) UpdateAge(user *entity.UpdateUser) string {
	for id, _ := range r.usersById {
		i := stringToInt(user.Target)
		if id == i {
			c := r.usersById[id]
			i2 := stringToInt(user.NewAge)
			c.Age = i2
		}
	}
	return "возраст пользователя успешно обновлен"
}

func stringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("err")
	}
	return i
}
