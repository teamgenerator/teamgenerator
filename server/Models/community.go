////
// author: Nico Alimin (nico@hackcapital.com)
// date: Sunday, 5th August 2018 1:24:59 am
// lastModifiedBy: Nico Alimin (nico@hackcapital.com)
// lastModifiedTime: Sunday, 5th August 2018 1:30:28 am
//
// DESCRIPTION
//
// copyright (c) 2018 Nico Alimin
////

package main

import (
	"database/sql"
	"errors"
)

type community struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location`
}

func (p *community) getProduct(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (p *community) updateProduct(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (p *community) deleteProduct(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (p *community) createProduct(db *sql.DB) error {
	return errors.New("Not implemented")
}

func getProducts(db *sql.DB, start, count int) ([]community, error) {
	return nil, errors.New("Not implemented")
}
