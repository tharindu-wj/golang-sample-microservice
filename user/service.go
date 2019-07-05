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

	subRouter.HandleFunc("/users", GetUsers).Methods("GET")
	subRouter.HandleFunc("/user/create", CreateUser).Methods("POST")
	subRouter.HandleFunc("/user/{id}", GetUser).Methods("GET")

	log.Fatal(http.ListenAndServe(":5000", router))
}

//get all users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	result := db.FindAll("users")

	user := model.User{}
	users := []model.User{}

	for result.Next() {
		var id int
		var name, email, created_at string
		err := result.Scan(&id, &name, &email, &created_at)
		if err != nil {
			panic(err.Error())
		}
		user.Id = id
		user.Name = name
		user.Email = email
		user.Created = created_at

		users = append(users, user)
	}

	json.NewEncoder(w).Encode(users)

}

//create new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	items := map[string]string{
		"name":    r.Form.Get("name"),
		"email":   r.Form.Get("email"),
		"company": r.Form.Get("company"),
	}

	result := db.Save("users", items)

	if (result) {
		json.NewEncoder(w).Encode("New User Created")
	}
}

//get single user
func GetUser(w http.ResponseWriter, r *http.Request) {
	userID := mux.Vars(r)["id"]

	result := db.FindBy("users", userID)

	user := model.User{}

	for result.Next() {
		var id int
		var name, email, created_at string
		err := result.Scan(&id, &name, &email, &created_at)
		if err != nil {
			panic(err.Error())
		}
		user.Id = id
		user.Name = name
		user.Email = email
		user.Created = created_at
	}

	json.NewEncoder(w).Encode(user)
}
