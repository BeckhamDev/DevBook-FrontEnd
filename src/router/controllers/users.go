package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"webapp/src/config"
	"webapp/src/responser"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	user, err := json.Marshal(map[string]string{
		"name": r.FormValue("name"),
		"email": r.FormValue("email"),
		"nick": r.FormValue("nick"),
		"password": r.FormValue("password"),
	})
	if err != nil {
		responser.JSON(w, http.StatusBadRequest, responser.Error{Erro: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/users", config.ApiURL)
	response, err := http.Post(url, "application/json", bytes.NewBuffer(user))

	if err != nil {
		responser.JSON(w, http.StatusInternalServerError, responser.Error{Erro: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responser.ErrorSanitize(w, response)
		return
	}

	responser.JSON(w, response.StatusCode, "success")
}