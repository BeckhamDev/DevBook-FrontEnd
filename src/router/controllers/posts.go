package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"webapp/src/config"
	"webapp/src/requests"
	"webapp/src/responser"

	"github.com/gorilla/mux"
)

// CreatePost calls an API method to create a new post on database
func CreatePost(w http.ResponseWriter, r *http.Request){
	r.ParseForm()

	post, err := json.Marshal(map[string]string{
		"title": r.FormValue("title"),
		"content": r.FormValue("content"),
	})

	if err != nil {
		responser.JSON(w, http.StatusBadRequest, responser.Error{Erro: err.Error()})
		return
	}
	fmt.Println(1)

	url := fmt.Sprintf("%s/Posts", config.ApiURL)
	response, err := requests.MakeAuthenticatedRequest(r, http.MethodPost, url, bytes.NewBuffer(post))

	if err != nil {
		responser.JSON(w, http.StatusInternalServerError, responser.Error{Erro: err.Error()})
		return
	}
	fmt.Println(2)
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		fmt.Println(5)
		responser.ErrorSanitize(w, response)
		return
	}
	fmt.Println(3)
	
	fmt.Println(4)
	responser.JSON(w, response.StatusCode, "success")
}

func LikePost(w http.ResponseWriter, r *http.Request){
	param := mux.Vars(r)

	postID, err := strconv.ParseUint(param["postID"], 10, 64)
	if err != nil {
		responser.JSON(w, http.StatusBadRequest, responser.Error{Erro: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/Posts/%d/Like", config.ApiURL, postID)
	response, err := requests.MakeAuthenticatedRequest(r, http.MethodPost, url, nil)
	if err != nil {
		responser.JSON(w, http.StatusInternalServerError, responser.Error{Erro: err.Error()})
		return
	}

	if response.StatusCode >= 400 {
		responser.ErrorSanitize(w, response)
		return
	}

	responser.JSON(w, response.StatusCode, "success")
}

func UnlikePost(w http.ResponseWriter, r *http.Request){
	param := mux.Vars(r)

	postID, err := strconv.ParseUint(param["postID"], 10, 64)
	if err != nil {
		responser.JSON(w, http.StatusBadRequest, responser.Error{Erro: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/Posts/%d/Unlike", config.ApiURL, postID)
	response, err := requests.MakeAuthenticatedRequest(r, http.MethodPost, url, nil)
	if err != nil {
		responser.JSON(w, http.StatusInternalServerError, responser.Error{Erro: err.Error()})
		return
	}

	if response.StatusCode >= 400 {
		responser.ErrorSanitize(w, response)
		return
	}

	responser.JSON(w, response.StatusCode, "success")
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)

	postID, err := strconv.ParseUint(param["postID"], 10, 64)
	if err != nil {
		responser.JSON(w, http.StatusBadRequest, responser.Error{Erro: err.Error()})
		return
	}

	r.ParseForm()
	post, err := json.Marshal(map[string]string{
		"title": r.FormValue("title"),
		"content": r.FormValue("content"),
	})

	if err != nil {
		responser.JSON(w, http.StatusBadRequest, responser.Error{Erro: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/Posts/%d", config.ApiURL, postID)
	response, err := requests.MakeAuthenticatedRequest(r, http.MethodPut, url, bytes.NewBuffer(post))
	if err != nil {
		responser.JSON(w, http.StatusInternalServerError, responser.Error{Erro: err.Error()})
		return
	}

	if response.StatusCode >= 400 {
		responser.ErrorSanitize(w, response)
		return
	}

	responser.JSON(w, response.StatusCode, "success")
}