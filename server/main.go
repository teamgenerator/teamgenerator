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
	"fmt"
	"log"
	"os"
	"strconv"
)

var (
	pgUser     = os.Getenv("PG_USER")
	pgPassword = os.Getenv("PG_PASSWORD")
	pgDatabase = os.Getenv("PG_DATABASE")
	pgHost     = os.Getenv("PG_HOST")
	pgPort     = os.Getenv("PG_PORT")
	port       = 3030
)

func main() {
	a := App{}

	a.Initialize(
		pgHost,
		pgPort,
		pgUser,
		pgPassword,
		pgDatabase,
	)
	fmt.Printf("Database initialized at %s for database %s\n", pgPort, pgDatabase)

	if _, err := a.DB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Initializing app\n")
	a.Run(":" + strconv.Itoa(port))
}

const tableCreationQuery = `CREATE TABLE IF NOT EXISTS community
(
id SERIAL,
name TEXT NOT NULL,
location TEXT NOT NULL,
CONSTRAINT community_pkey PRIMARY KEY (id)
)`
