package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"../shared/db"
	"../shared/model"
	"encoding/json"
)



//program starts from here
func main() {
	handleRequests()
}

//http request handler developed using mux package
func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	subRouter := router.PathPrefix("/api/v1").Subrouter()

	subRouter.HandleFunc("/posts", GetPosts).Methods("GET")
	subRouter.HandleFunc("/post/create", CreatePost).Methods("POST")
	subRouter.HandleFunc("/post/{id}", GetPost).Methods("GET")
	subRouter.HandleFunc("/post/delete/{id}", RemovePost).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", router))
}

//get all posts
func GetPosts(w http.ResponseWriter, r *http.Request) {
	result := db.FindAll("posts")
	post := model.Post{}
	posts := []model.Post{}
	for result.Next() {
		var id, user, category int
		var title, body, created_at string
		err := result.Scan(&id, &title, &body, &user, &category, &created_at)
		if err != nil {
			panic(err.Error())
		}
		post.Id = id
		post.Title = title
		post.Body = body
		post.User = user
		post.Category = category
		post.Created = created_at
		posts = append(posts, post)
	}

	json.NewEncoder(w).Encode(posts)

}

//get single post
func GetPost(w http.ResponseWriter, r *http.Request) {
	postID := mux.Vars(r)["id"]

	result := db.FindBy("posts", postID)

	post := model.Post{}

	for result.Next() {
		var id, user, category int
		var title, body, created_at string
		err := result.Scan(&id, &title, &body, &user, &category, &created_at)
		if err != nil {
			panic(err.Error())
		}
		post.Id = id
		post.Title = title
		post.Body = body
		post.User = user
		post.Category = category
		post.Created = created_at
	}

	json.NewEncoder(w).Encode(post)
}

//save post
func CreatePost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	items := map[string]string{
		"title": r.Form.Get("title"),
		"body":  r.Form.Get("body"),
		"user":  r.Form.Get("user"),
	}

	result := db.Save("posts", items)

	if (result) {
		json.NewEncoder(w).Encode("New Post Created")
	}
}

//save post
func RemovePost(w http.ResponseWriter, r *http.Request) {
	postID := mux.Vars(r)["id"]

	result := db.Remove("posts", postID)

	if (result) {
		json.NewEncoder(w).Encode("Post Deleted")
	}
}