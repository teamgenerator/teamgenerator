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
	pgUser     = os.Getenv("PG_USER")
	pgPassword = os.Getenv("PG_PASSWORD")
	pgDatabase = os.Getenv("PG_DATABASE")
	port       = 3030
)

func main() {
	a := App{}
	a.Initialize(
		pgUser,
		pgPassword,
		pgDatabase,
	)

	if _, err := a.DB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}

	a.Run(":8080")
	log.Printf("The Go Api server is listening on port 3030")
}

const tableCreationQuery = `CREATE TABLE IF NOT EXISTS community
(
id SERIAL,
name TEXT NOT NULL,
location TEXT NOT NULL,
CONSTRAINT community_pkey PRIMARY KEY (id)
)`
