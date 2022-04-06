package handlers

import (
	"backend-sv/connection"
	"backend-sv/structs"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome!")
}

func CreateArticle(w http.ResponseWriter, r *http.Request) {
	payloads, _ := ioutil.ReadAll(r.Body)

	var article structs.Posts
	article.Created_date = time.Now()
	article.Updated_date = time.Now()

	json.Unmarshal(payloads, &article)

	if article.Title == "" || article.Content == "" || article.Category == "" || article.Status == "" {
		http.Error(w, "Please enter the required fields", http.StatusBadRequest)
	} else {
		connection.DB.Create(&article)

		res := structs.Result{Code: 200, Data: article, Message: "Success create article"}

		result, err := json.Marshal(res)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(result)
	}
}

func GetArticles(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	take := r.URL.Query().Get("take")
	if len(page) == 0 || len(take) == 0 {
		take = "100"
		page = "0"
	}

	articles := []structs.Posts{}

	connection.DB.
		Limit(take).
		Offset(page).
		Order("id").
		Find(&articles)

	res := structs.Result{Code: 200, Data: articles, Message: "Success get articles, page: " + page + ", take: " + take}
	results, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(results)
}

func GetArticleById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	articleID := vars["id"]

	var article structs.Posts

	connection.DB.Where("id = ?", articleID).First(&article)

	res := structs.Result{Code: 200, Data: article, Message: "Success get article"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func UpdateArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	articleID := vars["id"]

	payloads, _ := ioutil.ReadAll(r.Body)

	var article structs.Posts
	article.Updated_date = time.Now()
	json.Unmarshal(payloads, &article)

	connection.DB.First(&article, articleID)

	json.Unmarshal(payloads, &article)

	article.Updated_date = time.Now()

	connection.DB.Save(&article)

	res := structs.Result{Code: 200, Data: &article, Message: "Success update article"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	articleID := vars["id"]

	var article structs.Posts
	connection.DB.Where("id = ?", articleID).First(&article)

	if article.ID == 0 {
		http.Error(w, "Article not found", http.StatusBadRequest)
	} else {
		connection.DB.Delete(&article)
		connection.DB.Where("id = ?", articleID).Delete(&article)
		res := structs.Result{Code: 200, Data: article, Message: "Success delete article"}
		result, err := json.Marshal(res)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(result)
	}
}
