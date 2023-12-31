package cookies

import (
	"net/http"
	"time"
	"webapp/src/config"

	"github.com/gorilla/securecookie"
)

var s *securecookie.SecureCookie 

func Config() {
	s = securecookie.New(config.HashKey, config.BlockKey)
}

func Save(w http.ResponseWriter, ID, Token string) error {
	AuthData := map[string]string{
		"id": ID,
		"token": Token,
	}

	dataEncoded, err := s.Encode("AuthData", AuthData)
	if err != nil {
		return err
	}

	http.SetCookie(w, &http.Cookie{
		Name: "AuthData",
		Value: dataEncoded,
		Path: "/",
		HttpOnly: true,
	})

	return nil
}

func ReadCookies(r *http.Request) (map[string]string, error){
	cookie, err := r.Cookie("AuthData")
	if err != nil {
		return nil, err
	}

	values := make(map[string]string)
	if err = s.Decode("AuthData", cookie.Value, &values); err != nil {
		return nil, err
	}

	return values, nil
}

func Delete(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name: "AuthData",
		Value: "",
		Path: "/",
		HttpOnly: true,
		Expires: time.Unix(0,0),
	})
}