////
// author: Nico Alimin (nico@hackcapital.com)
// date: Sunday, 5th August 2018 12:48:35 am
// lastModifiedBy: Nico Alimin (nico@hackcapital.com)
// lastModifiedTime: Sunday, 5th August 2018 12:52:42 am
//
// DESCRIPTION
//
// copyright (c) 2018 Nico Alimin
////

package main

import (
	"log"

	// External Dependency
	"github.com/gorilla/mux"
)

var (
	port = 3030
)

func main() {
	main := mux.NewRouter()

	routerAPI := main.PathPrefix("/api").Subrouter()
	routerV1 := main.PathPrefix("/v1").Subrouter()

	log.Printf("The Go Api server is listening on port 3030")
}
