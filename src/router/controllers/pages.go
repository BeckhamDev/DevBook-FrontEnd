package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/models"
	"webapp/src/requests"
	"webapp/src/responser"
	"webapp/src/utils"

	"github.com/gorilla/mux"
)

func LoadLoginPage(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.ReadCookies(r)
	if cookie["token"] != "" {
		http.Redirect(w, r,"/home", http.StatusFound)
		return
	}
	
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

func LoadPostEditPage(w http.ResponseWriter, r *http.Request){
	param := mux.Vars(r)

	postID, err := strconv.ParseUint(param["postID"], 10, 64)
	if err != nil {
		responser.JSON(w, http.StatusBadRequest, responser.Error{Erro: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/Posts/%d", config.ApiURL, postID)
	response, err := requests.MakeAuthenticatedRequest(r, http.MethodGet, url, nil)
	if err != nil {
		responser.JSON(w, http.StatusInternalServerError, responser.Error{Erro: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responser.ErrorSanitize(w, response)
		return
	}

	var post models.Post
	if err = json.NewDecoder(response.Body).Decode(&post);err != nil{
		responser.JSON(w, http.StatusUnprocessableEntity, responser.Error{Erro: err.Error()})
		return
	}

	utils.ExecTemplate(w, "edit_post.html", post)
}

func LoadUsersPage(w http.ResponseWriter, r *http.Request){
	nameOrNick := strings.ToLower(r.URL.Query().Get("user"))
	url := fmt.Sprintf("%s/users?user=%s",config.ApiURL ,nameOrNick)

	response, err := requests.MakeAuthenticatedRequest(r, http.MethodGet, url, nil)
	if err != nil {
		responser.JSON(w, http.StatusInternalServerError, responser.Error{Erro: "O erro foi aqui"})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responser.ErrorSanitize(w, response)
		return
	}

	var user []models.User
	if err := json.NewDecoder(response.Body).Decode(&user); err != nil {
		responser.JSON(w, http.StatusUnprocessableEntity, responser.Error{Erro: err.Error()})
		return
	}

	utils.ExecTemplate(w, "users.html", user)
}

func LoadUserProfile(w http.ResponseWriter, r *http.Request){
	param := mux.Vars(r)

	userID, err := strconv.ParseUint(param["userID"], 10, 64)
	if err != nil {
		responser.JSON(w, http.StatusBadRequest, responser.Error{Erro: err.Error()})
		return
	}

	user, err := models.SearchAllUserInfo(userID, r)
	if err != nil {
		responser.JSON(w, http.StatusInternalServerError, responser.Error{Erro: err.Error()})
		return
	}

	cookie, _ := cookies.ReadCookies(r)
	userLogged, _ := strconv.ParseUint(cookie["id"], 10, 64) 

	utils.ExecTemplate(w, "user.html", struct {
		User models.User
		UserLogged uint64
	}{
		User: user,
		UserLogged: userLogged,
	})

}