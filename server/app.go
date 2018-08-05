////
// author: Nico Alimin (nico@hackcapital.com)
// date: Sunday, 5th August 2018 1:24:25 am
// lastModifiedBy: Nico Alimin (nico@hackcapital.com)
// lastModifiedTime: Sunday, 5th August 2018 1:24:25 am
//
// DESCRIPTION
//
// copyright (c) 2018 Nico Alimin
////

package main

import (
	"database/sql"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Initialize(user, password, dbname string) {}

func (a *App) Run(addr string) {}
