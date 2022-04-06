package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func HandleReq() {
	log.Println("Start development server localhost:9999")

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", HomePage)
	myRouter.HandleFunc("/article", CreateArticle).Methods("OPTIONS", "POST")
	myRouter.HandleFunc("/articles/", GetArticles).Methods("OPTIONS", "GET")
	myRouter.HandleFunc("/article/{id}", GetArticleById).Methods("OPTIONS", "GET")
	myRouter.HandleFunc("/article/{id}", UpdateArticle).Methods("OPTIONS", "PATCH")
	myRouter.HandleFunc("/article/{id}", DeleteArticle).Methods("OPTIONS", "Delete")

	handler := cors.AllowAll().Handler(myRouter)
	log.Fatal(http.ListenAndServe(":8000", handler))
}
