package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

/*Напишите HTTP-сервис, который принимает входящие соединения с JSON-данными и обрабатывает их следующим образом:
    1. Сделайте обработчик создания пользователя. У пользователя должны быть следующие поля: имя, возраст и массив друзей.
Пользователя необходимо сохранять в мапу. Пример запроса:
Данный запрос должен возвращать ID пользователя и статус 201.
*/

type User struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Friends []*User `json:"friends"`
}

func (u *User) toString() string {
	return fmt.Sprintf("name is %s and age is %d and friends %v \n", u.Name, u.Age, u.Friends)
}
type service struct {
	i int
	store map [int] *User
}

func main()  {
	mux := http.NewServeMux()
	srv := service{i: int(1), store: map[int]*User{}}
	mux.HandleFunc("/create", srv.Create)
//	mux.HandleFunc("/get", srv.GetAll	)
	http.ListenAndServe("localhost:8080", mux)
}

func (s *service)  Create (w http.ResponseWriter, r *http.Request)  {
	if r.Method == "POST" {
		content, err := ioutil.ReadAll(r.Body) // берем и считываем body (информацию с rest запроса) и заносим в переменную content (срез байтов)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError) //статус серверная ошибка
			w.Write([]byte(err.Error()))
			return
		}
		//fmt.Println(content,s.i, "тый контент")
		defer r.Body.Close()                                //отложенное закрытие считывателя
		var u User                                          //создаем n- ого пользователя
		if err := json.Unmarshal(content, &u); err != nil { /*благодаря Unmarshal полученная информация в байтах (content) заносится
			в данные созданного пользователя по соответствующим полям*/
			w.WriteHeader(http.StatusInternalServerError) //статус серверная ошибка
			w.Write([]byte(err.Error()))
			return
		}
		s.store[s.i] = &u // записывает в мапу с ключом = имени из данных пользователя и присваивает данному ключу полученные данные

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("пользователь " + u.Name + " зарегистрирован под номером " + strconv.Itoa(s.i))) // отправляем в POST информацию, что всё прошло ок
		//fmt.Println(u)
		s.i++
		//fmt.Println(s)
	}
	w.WriteHeader(http.StatusBadRequest)
}
/*
func (s  *service)GetAll(w http.ResponseWriter, r *http.Request)  {
	if r.Method == "GET"{
		response := ""
		for _, user := range s.store { // проходит итерацию по всей мапе и выводит информацию от сервера по позициям из мапы
			response += user.toString()
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(response))
		return
	}
	w.WriteHeader(http.StatusBadRequest)
}*/