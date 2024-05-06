package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type connDB struct {
	DB *sqlx.DB
}

func (c *connDB) Close() {
	c.DB.Close()
}

func NewConnectDB() (*connDB, error) {
	db, err := sqlx.Open("sqlite3", "data.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	stmt, err := db.Prepare("CREATE TABLE IF NOT EXISTS cotacao (" +
		"id INTEGER PRIMARY KEY, " +
		"code TEXT,codein TEXT, " +
		"name TEXT, " +
		"high TEXT, " +
		"low TEXT, " +
		"varbid TEXT, " +
		"pctchange TEXT, " +
		"bid TEXT, " +
		"ask TEXT, " +
		"timestamp TEXT, " +
		"createdate TEXT)")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return nil, err
	}

	return &connDB{DB: db}, nil
}
