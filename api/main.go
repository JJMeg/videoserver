package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func RegisterHandlers() *httprouter.Router{
	router := httprouter.New()
	//user
	router.POST("/user",CreateUser)
	router.POST("/user/:user_name",Login)
	//router.GET("/user/:username",);
	//
	return router
}

func main(){
	r := RegisterHandlers()
	http.ListenAndServe(":8000",r)//阻塞在此
}

