package controller

import (
	"Assignment/database"
	"Assignment/model"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// Gets All the Blog Posts Contains in the Database
func GetBlogPosts(w http.ResponseWriter, r *http.Request) {
	var blogs []model.BlogPosts
	database.Instance.Find(&blogs)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(blogs)
}

// Gets a Blog Post by id Contains in the Database
func GetBlogPostByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var blog model.BlogPosts
	database.Instance.First(&blog, id)
	if blog.Id == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Blog Not Found")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(blog)
}

// Creates a Blog Post in the Database
func CreateBlogPost(w http.ResponseWriter, r *http.Request) {
	var blog model.BlogPosts
	w.Header().Set("Content-Type", "application/json")
	json.NewDecoder(r.Body).Decode(&blog)
	tx := database.Instance.Create(&blog)
	if tx.Error != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(tx.Error.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Blog Post Created!")
}

// Updates a Blog Post in the Database
func UpdateBlogPost(w http.ResponseWriter, r *http.Request) {
	var blog model.BlogPosts
	json.NewDecoder(r.Body).Decode(&blog)
	if findBlog(blog.Id) {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Blog Not Found")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	tx := database.Instance.Save(&blog)
	if tx.Error != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(tx.Error.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Blog Post Updated!")
}

// Deletes a Blog Post in the Database
func DeleteBlogPost(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var blog model.BlogPosts
	database.Instance.First(&blog, id)
	if blog.Id == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Blog Not Found")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	tx := database.Instance.Delete(&blog, id)
	if tx.Error != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(tx.Error.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Blog Post Deleted!")
}

func findBlog(id int) bool {
	var blog model.BlogPosts
	database.Instance.First(&blog, id)
	return blog.Id == 0
}
