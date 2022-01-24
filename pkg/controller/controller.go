package controller

import (
	"encoding/json"
	"homework.30/pkg/entity"
	"homework.30/pkg/usecase"
	"net/http"
)

type Controller struct {
	usecase usecase.Usecase
}

func NewController(usecase usecase.Usecase) *Controller {
	return &Controller{
		usecase: usecase,
	}
}

func (c *Controller) CreateUser(w http.ResponseWriter, r *http.Request) {
	user := &entity.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusBadRequest)
	}
	id, err := c.usecase.CreateUser(user)
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusBadRequest)
	}
	result := map[string]int{"id": id}
	response, err := json.Marshal(result)
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}
