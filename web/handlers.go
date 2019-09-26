package main

import (
	"github.com/julienschmidt/httprouter"
	"html/template"
	"log"
	"net/http"
)

type HomePage struct {
	Name string
}

func homeHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	p := &HomePage{Name: "avenssi"}
	t, e := template.ParseFiles("./templates/home.html")
	if e != nil {
		log.Printf("Parsing templates home.html error:%s", e)
		return
	}
	//把模板执行转给浏览器
	t.Execute(w, p)
	return
}
