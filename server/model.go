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
)

type community struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
}

func (p *community) getCommunity(db *sql.DB) error {
	return db.QueryRow("SELECT id, name, Location FROM community WHERE id=$1",
		p.ID).Scan(&p.ID, &p.Name, &p.Location)
}

func (p *community) updateCommunity(db *sql.DB) error {
	_, err :=
		db.Exec("UPDATE community SET name=$1, location=$2 WHERE id=$3",
			p.Name, p.Location, p.ID)

	return err
}

func (p *community) deleteCommunity(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM community WHERE id=$1", p.ID)

	return err
}

func (p *community) createCommunity(db *sql.DB) error {
	err := db.QueryRow(
		"INSERT INTO community(name, location) VALUES($1, $2) RETURNING id",
		p.Name, p.Location).Scan(&p.ID)

	if err != nil {
		return err
	}

	return nil
}

func getCommunities(db *sql.DB, start, count int) ([]community, error) {
	rows, err := db.Query(
		"SELECT id, name, location FROM community LIMIT $1 OFFSET $2",
		count, start)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	communities := []community{}

	for rows.Next() {
		var p community
		if err := rows.Scan(&p.ID, &p.Name, &p.Location); err != nil {
			return nil, err
		}
		communities = append(communities, p)
	}

	return communities, nil
}
