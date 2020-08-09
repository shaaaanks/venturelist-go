package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World")
}

func getProjects(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	projects, err := database.FindAll()

	if err != nil {
		log.Fatalf("Error getting items from database: %v", err)
	}

	json.NewEncoder(w).Encode(projects)
}

func getProject(w http.ResponseWriter, r *http.Request) {
	projectID := mux.Vars(r)["id"]

	project, err := database.Find(projectID)
	if err != nil {
		log.Fatalf("Error getting item from database: %v", err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(project)
}

func createProject(w http.ResponseWriter, r *http.Request) {
	var project project

	request, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error reading Project from request: %v", err)
	}

	json.Unmarshal(request, &project)

	database.Create(project)

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(project)
}

func createRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", index).Methods("GET")
	router.HandleFunc("/project", createProject).Methods("POST")
	router.HandleFunc("/projects", getProjects).Methods("GET")
	router.HandleFunc("/project/{id}", getProject).Methods("GET")

	return router
}
