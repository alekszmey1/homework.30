package main

import (
	//"encoding/json"
	//"fmt"
	"homework.30/pkg/service"
	"homework.30/pkg/user"
	//"io/ioutil"
	"net/http"
)

/*Напишите HTTP-сервис, который принимает входящие соединения с JSON-данными и обрабатывает их следующим образом:
    1. Сделайте обработчик создания пользователя. У пользователя должны быть следующие поля: имя, возраст и массив друзей.
Пользователя необходимо сохранять в мапу. Пример запроса:
Данный запрос должен возвращать ID пользователя и статус 201.
*/



func main()  {
	mux := http.NewServeMux()
	srv := service.Service{Store: map[int]*user.User{}}
	mux.HandleFunc("/create", srv.Create)
//	mux.HandleFunc("/get", srv.GetAll	)
	http.ListenAndServe("localhost:8080", mux)
}


/*
func (s  *service)GetAll(w http.ResponseWriter, r *http.Request)  {
	if r.Method == "GET"{
		response := ""
		for _, user := range s.store { // проходит итерацию по всей мапе и выводит информацию от сервера по позициям из мапы
			response += User.ToString()
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(response))
		return
	}
	w.WriteHeader(http.StatusBadRequest)
}*/