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

	subRouter.HandleFunc("/companies", GetCompanies).Methods("GET")
	subRouter.HandleFunc("/company/{id}", GetCompany).Methods("GET")

	log.Fatal(http.ListenAndServe(":2000", router))
}

//get all companies
func GetCompanies(w http.ResponseWriter, r *http.Request) {
	result := db.FindAll("companies")
	company := model.Company{}
	companies := []model.Company{}
	for result.Next() {
		var id int
		var name, location, created_at string
		err := result.Scan(&id, &name, &location, &created_at)
		if err != nil {
			panic(err.Error())
		}
		company.Id = id
		company.Name = name
		company.Location = location
		company.Created = created_at
		companies = append(companies, company)
	}

	json.NewEncoder(w).Encode(companies)
}

//get single post
func GetCompany(w http.ResponseWriter, r *http.Request) {
	compID := mux.Vars(r)["id"]

	result := db.FindBy("companies", compID)

	company := model.Company{}

	for result.Next() {
		var id int
		var name, location, created_at string
		err := result.Scan(&id, &name, &location, &created_at)
		if err != nil {
			panic(err.Error())
		}
		company.Id = id
		company.Name = name
		company.Location = location
		company.Created = created_at
	}

	json.NewEncoder(w).Encode(company)
}
