package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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
type Service struct {
	Id int `json:"Id"`
	store map [int] *User
}
//пробую  добавлять карту в сервис через отдельный метод
func (s *Service) GetMap(u *User)  {
	for {
		s.Id++
		s.store[s.Id] = u
		fmt.Printf("Запись %v произведена \n", s.Id)
		return
	}
}
func main()  {
	mux := http.NewServeMux()
	srv := Service{store: map[int]*User{}}
	mux.HandleFunc("/create", srv.Create)
//	mux.HandleFunc("/get", srv.GetAll	)
	http.ListenAndServe("localhost:8080", mux)
}

func (s *Service)  Create (w http.ResponseWriter, r *http.Request)  {
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
		//добавил метод обработки добавления структуры пользователя в мапу
		s.GetMap(&u)
		w.WriteHeader(http.StatusCreated)
		//w.Write([]byte("пользователь " + u.Name + " зарегистрирован под номером " + strconv.Itoa(s.Id))) // отправляем в POST информацию, что всё прошло ок
		//fmt.Println(u)
		//s.Id++
		//fmt.Println(s)
		data, err := json.Marshal(s.Id)
		if err != nil {
			fmt.Println(err)
			return
		}
		w.Write(data)
	}
	//w.WriteHeader(http.StatusBadRequest)// убрал, т.к. писала 2022/01/07 11:11:00 http: superfluous response.WriteHeader call from main.(*service).Create (main.go:75)
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