package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/models"
	"webapp/src/responser"
)

func Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	user, err := json.Marshal(map[string]string{
		"email": r.FormValue("email"),
		"password": r.FormValue("password"),
	})
	if err != nil {
		responser.JSON(w, http.StatusBadRequest, responser.Error{Erro: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/login", config.ApiURL)
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

	var authData models.AuthData
	if err := json.NewDecoder(response.Body).Decode(&authData); err != nil {
		responser.JSON(w, http.StatusUnprocessableEntity, responser.Error{Erro: err.Error()})
		return
	}

	if err := cookies.Save(w, authData.ID, authData.Token); err != nil {
		responser.JSON(w, http.StatusUnprocessableEntity, responser.Error{Erro: err.Error()})
		return
	}
	
	responser.JSON(w, response.StatusCode, "success")
}