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
	"os"
)

var (
	PG_USER     = os.Getenv("PG_USER")
	PG_PASSWORD = os.Getenv("PG_PASSWORD")
	PG_DATABASE = os.Getenv("PG_DATABASE")
	port        = 3030
)

func main() {
	a := App{}
	a.Initialize(
		PG_USER,
		PG_PASSWORD,
		PG_DATABASE,
	)

	a.Run(":8080")
	log.Printf("The Go Api server is listening on port 3030")
}
