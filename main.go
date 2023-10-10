package main

import (
	"Assignment/config"
	"Assignment/controller"
	"Assignment/database"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var DB *gorm.DB

func main() {
	fmt.Println("Assignment")

	config.LoadConfig()

	// Database Connection
	database.Connect(config.Configuration.ConnectionString)
	database.Migrate()

	//Initialize router
	router := mux.NewRouter().StrictSlash(true)
	InitialzeRoutes(router)

	// Starting Server
	log.Printf(fmt.Sprintf("Starting Server on port %s", config.Configuration.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", config.Configuration.Port), router))
}

func InitialzeRoutes(router *mux.Router) {
	router.HandleFunc("/api/getPosts", controller.GetBlogPosts).Methods("GET")           //Gets All Blog-Posts
	router.HandleFunc("/api/getPosts/{id}", controller.GetBlogPostByID).Methods("GET")   //Gets an Blog-Post by ID
	router.HandleFunc("/api/createPost", controller.CreateBlogPost).Methods("POST")      //Creates a Blog-Post
	router.HandleFunc("/api/updatePost", controller.UpdateBlogPost).Methods("PUT")  //Updates a Blog-Post
	router.HandleFunc("/api/deletePost/{id}", controller.DeleteBlogPost).Methods("DELETE") //Deletes a Blog-Post
}
