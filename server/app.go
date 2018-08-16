////
// author: Nico Alimin (nicoalimin@gmail.com)
// date: Sunday, 5th August 2018 1:24:25 am
// lastModifiedBy: Nico Alimin (nicoalimin@gmail.com)
// lastModifiedTime: Sunday, 5th August 2018 1:24:25 am
//
// DESCRIPTION
//
// copyright (c) 2018 Nico Alimin
////

package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// The App object that consists of:
// - Router
// - Database
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

// Initialize the database and router
func (a *App) Initialize(host, port, user, password, dbname string) {
	connectionString :=
		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	fmt.Printf("Connection String is: %s\n", connectionString)

	var err error
	a.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	//
	// Creates the router and subrouters
	//

	// Main router
	main := mux.NewRouter()

	// API subrouter
	routerAPI := main.PathPrefix("/api").Subrouter()

	// Version subrouter
	routerV1 := routerAPI.PathPrefix("/v1").Subrouter()

	a.Router = routerV1

	a.initializeRoutes()
}

// Run the server at designated port
func (a *App) Run(addr string) {
	fmt.Printf("The Go Api server is listening on port %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/communities", a.getCommunities).Methods("GET")
	a.Router.HandleFunc("/community", a.createCommunity).Methods("POST")
	a.Router.HandleFunc("/community/{id:[0-9]+}", a.getCommunity).Methods("GET")
	a.Router.HandleFunc("/community/{id:[0-9]+}", a.updateCommunity).Methods("PUT")
	a.Router.HandleFunc("/community/{id:[0-9]+}", a.deleteCommunity).Methods("DELETE")
}

func (a *App) getCommunities(w http.ResponseWriter, r *http.Request) {
	count, _ := strconv.Atoi(r.FormValue("count"))
	start, _ := strconv.Atoi(r.FormValue("start"))

	if count > 10 || count < 1 {
		count = 10
	}
	if start < 0 {
		start = 0
	}

	community, err := getCommunities(a.DB, start, count)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, community)
}

func (a *App) getCommunity(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Community ID")
		return
	}

	p := community{ID: id}
	if err := p.getCommunity(a.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Community not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, p)
}

func (a *App) createCommunity(w http.ResponseWriter, r *http.Request) {
	var p community
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := p.createCommunity(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, p)
}

func (a *App) updateCommunity(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid community ID")
		return
	}

	var p community
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()
	p.ID = id

	if err := p.updateCommunity(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, p)
}

func (a *App) deleteCommunity(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Community ID")
		return
	}

	p := community{ID: id}
	if err := p.deleteCommunity(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}