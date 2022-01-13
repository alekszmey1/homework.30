package service

import (
	"homework.30/pkg/repository"
	"encoding/json"
	"fmt"
	"homework.30/pkg/user"
	"io/ioutil"
	"net/http"
)


func (s *repository)  Create (w http.ResponseWriter, r *http.Request)  {
	if r.Method == "POST" {
		content, err := ioutil.ReadAll(r.Body) // берем и считываем body (информацию с rest запроса) и заносим в переменную content (срез байтов)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError) //статус серверная ошибка
			w.Write([]byte(err.Error()))
			return
		}
		defer r.Body.Close()                                //отложенное закрытие считывателя
		var u user.User                                        //создаем n- ого пользователя
		if err := json.Unmarshal(content, &u); err != nil { /*благодаря Unmarshal полученная информация в байтах (content) заносится
			в данные созданного пользователя по соответствующим полям*/
			w.WriteHeader(http.StatusInternalServerError) //статус серверная ошибка
			w.Write([]byte(err.Error()))
			return
		}
		s.repository.CreateUser(&u)
		w.WriteHeader(http.StatusCreated)
		t := UserId{s.Id}
		data, err := json.Marshal(t)
		if err != nil {
			fmt.Println(err)
			return
		}
		w.Write(data)
	}
	//w.WriteHeader(http.StatusBadRequest)// убрал, т.к. писала 2022/01/07 11:11:00 http: superfluous response.WriteHeader call from main.(*service).Create (main.go:75)
}