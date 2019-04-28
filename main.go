package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"videoserver/api"
)

func RegisterHandlers() *httprouter.Router{
	router := httprouter.New()
	//user
	router.POST("/user",api.CreateUser)
	router.POST("/user/:user_name",api.Login)

	//
	return router
}

func main(){
	r := RegisterHandlers()
	http.ListenAndServe(":8000",r)//阻塞在此
}

