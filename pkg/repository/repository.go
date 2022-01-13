package repository

import (
	"fmt"
	"homework.30/pkg/user"
)

type repository struct {
	Id        int
	usersById map [int] * user.User
}
type UserId struct {
	UserId int
}

//пробую  добавлять карту в сервис через отдельный метод
func (r *repository) CreateUser (user *user.User)  {
	for {
		r.Id++
		user.Id = r.Id
		r.usersById[r.Id] = user
		fmt.Printf("Запись %v произведена \n", r.Id)
		return
	}
}

func (r *repository)MakeFriends(a, b int,user, user2 *user.User)  {
	user.Friends[a] = append(friends)
}

func (r * repository)DeleteUser(id int){

}