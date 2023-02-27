package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"

	"github.com/api"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/user", controller.GetAll).Methods("GET")
	router.HandleFunc("/user", controller.Insert).Methods("POST")
	router.HandleFunc("/user", controller.Update).Methods("PUT")
	router.HandleFunc("/user", controller.Delete).Methods("DELETE")
	router.HandleFunc("/", controller.MainPage).Methods("GET")
	http.Handle("/", router)
	fmt.Println("Connected to port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}


