package user

import "fmt"

type User struct {
	Name string `json:"name"`
	Age int         `json:"age"`
	Friends []*User `json:"friends"`
}

func (u *User) ToString() string {
	return fmt.Sprintf("name is %s and age is %d and friends %v \n", u.Name, u.Age, u.Friends)
}