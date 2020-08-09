package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func getProjects(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	projects, err := database.FindAll()

	if err != nil {
		fmt.Fprintf(w, "Error getting items from database: %v", err)
		return
	}

	json.NewEncoder(w).Encode(projects)
}

func getProject(w http.ResponseWriter, r *http.Request) {
	projectID := mux.Vars(r)["id"]

	project, err := database.Find(projectID)
	if err != nil {
		fmt.Fprintf(w, "Error getting item from database: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(project)
}

func createProject(w http.ResponseWriter, r *http.Request) {
	var project project

	request, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error reading Project from request: %v", err)
		return
	}

	json.Unmarshal(request, &project)

	err = validate(project)
	if err != nil {
		fmt.Fprintf(w, "Validation error: %v", err)
		return
	}

	database.Create(project)

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(project)
}

func updateProject(w http.ResponseWriter, r *http.Request) {
	projectID := mux.Vars(r)["id"]
	var project project

	request, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error reading Project from request: %v", err)
		return
	}

	json.Unmarshal(request, &project)

	err = validate(project)
	if err != nil {
		fmt.Fprintf(w, "Validation error: %v", err)
		return
	}

	err = database.Update(projectID, project)
	if err != nil {
		fmt.Fprintf(w, "Error updating item in database: %v", err)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(project)
}

func deleteProject(w http.ResponseWriter, r *http.Request) {
	projectID := mux.Vars(r)["id"]

	err := database.Delete(projectID)
	if err != nil {
		fmt.Fprintf(w, "Error deleting item from database: %v", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "The project with the ID %v has been deleted successfully", projectID)
}

func createRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/project", createProject).Methods("POST")
	router.HandleFunc("/projects", getProjects).Methods("GET")
	router.HandleFunc("/project/{id}", getProject).Methods("GET")
	router.HandleFunc("/project/{id}", updateProject).Methods("PATCH")
	router.HandleFunc("/project/{id}", deleteProject).Methods("DELETE")

	return router
}
