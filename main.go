package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/router"
	"webapp/src/utils"
)

func main() {
	defer fmt.Println("Rodando Web Application")
	config.Load()
	cookies.Config()
	utils.LoadTemplates()
	r := router.Router()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}