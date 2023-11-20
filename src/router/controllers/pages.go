package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/models"
	"webapp/src/requests"
	"webapp/src/responser"
	"webapp/src/utils"
)

func LoadLoginPage(w http.ResponseWriter, r *http.Request) {
	utils.ExecTemplate(w, "login.html", nil)
}

func LoadCreateUserPage(w http.ResponseWriter, r *http.Request) {
	utils.ExecTemplate(w, "create_user.html", nil)
}

func LoadHomePage(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/Posts", config.ApiURL)
	response, err:= requests.MakeAuthenticatedRequest(r, http.MethodGet, url, nil)

	if err != nil {
		responser.JSON(w, http.StatusInternalServerError, responser.Error{Erro: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responser.ErrorSanitize(w, response)
		return
	}
	
	var posts []models.Post
	if err := json.NewDecoder(response.Body).Decode(&posts);err != nil {
		responser.JSON(w, http.StatusUnprocessableEntity, responser.Error{Erro: err.Error()})
		return
	}

	cookie, _ := cookies.ReadCookies(r)
	userID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	utils.ExecTemplate(w, "home.html", struct{
		Posts []models.Post
		UserID uint64 
	}{
		Posts: posts,
		UserID: userID,
	})
}