package controllers

import "net/http"


func CreateUser(w http.ResponseWriter, r *http.Request) {
   	w.Write([]byte("welcome from user controller"))
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
   	w.Write([]byte("welcome from user controller"))
}