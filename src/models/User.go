package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
	"webapp/src/config"
	"webapp/src/requests"
)

type User struct {
	ID        uint64  `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Nick      string `json:"nick"`
	CreatedAt time.Time `json:"created_at"`
	Followers []User `json:"followers"`
	Following []User `json:"following"`
	Posts 	  []Post `json:"posts"`
}

func SearchAllUserInfo(userID uint64, r *http.Request) (User, error){
	userChannel := make(chan User)
	followersChannel := make(chan []User)
	followingChannel := make(chan []User)
	postChannel := make(chan []Post)

	go SearchUserInfo(userChannel, userID, r)
	go SearchFollowers(followersChannel, userID, r)
	go SearchFollowing(followingChannel, userID, r)
	go SearchPosts(postChannel, userID, r)

	var (
		user User
		followers []User
		following []User
		posts []Post
	)

	for i:=0; i<4; i++ {
		select{
		case userLoaded := <-userChannel:
			if userLoaded.ID == 0 {
				return User{}, errors.New("error loading user")
			}

			user = userLoaded

		case followersLoaded := <-followersChannel:
			if followersLoaded == nil {
				return User{}, errors.New("error loading followers")
			}

			followers = followersLoaded

		case followingLoaded := <-followingChannel:
			if followingLoaded == nil {
				return User{}, errors.New("error loading users you follow")
			}

			following = followingLoaded

		case postsLoaded := <-postChannel:
			if postsLoaded == nil {
				return User{}, errors.New("error loading posts")
			}

			posts = postsLoaded
		}
	}
	user.Followers = followers
	user.Following = following
	user.Posts = posts

	return user, nil
}

func SearchUserInfo(channel chan<- User, userID uint64, r *http.Request){
	url := fmt.Sprintf("%s/users/%d", config.ApiURL, userID)
	response, err := requests.MakeAuthenticatedRequest(r, http.MethodGet, url, nil)
	if err != nil {
		channel <- User{}
		return
	}
	defer response.Body.Close()

	var user User
	if err := json.NewDecoder(response.Body).Decode(&user); err != nil {
		channel <- User{}
		return
	}

	channel <- user
}

func SearchFollowers(channel chan<- []User, userID uint64, r *http.Request){
	url := fmt.Sprintf("%s/users/%d/Followers", config.ApiURL, userID)
	response, err := requests.MakeAuthenticatedRequest(r, http.MethodGet, url, nil)
	if err != nil {
		channel <- nil
		return
	}
	defer response.Body.Close()

	var followers []User
	if err := json.NewDecoder(response.Body).Decode(&followers); err != nil {
		channel <- nil
		return
	}

	channel <- followers
}

func SearchFollowing(channel chan<- []User, userID uint64, r *http.Request){
	url := fmt.Sprintf("%s/users/%d/Following", config.ApiURL, userID)
	response, err := requests.MakeAuthenticatedRequest(r, http.MethodGet, url, nil)
	if err != nil {
		channel <- nil
		return
	}
	defer response.Body.Close()

	var following []User
	if err := json.NewDecoder(response.Body).Decode(&following); err != nil {
		channel <- nil
		return
	}

	channel <- following
}

func SearchPosts(channel chan<- []Post, userID uint64, r *http.Request){
	url := fmt.Sprintf("%s/Users/%d/Posts", config.ApiURL, userID)
	response, err := requests.MakeAuthenticatedRequest(r, http.MethodGet, url, nil)
	if err != nil {
		channel <- nil
		return
	}
	defer response.Body.Close()

	var posts []Post
	if err := json.NewDecoder(response.Body).Decode(&posts); err != nil {
		channel <- nil
		return
	}

	channel <- posts
}