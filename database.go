package main

import (
	"database/sql"
	"fmt"
	"net/url"
)

func OpenDB(host, username, token string) (*sql.DB, error) {
	url := &url.URL{
		Scheme: "sqlserver",
		User:   url.UserPassword(username, token),
		Host:   host,
	}

	db, err := sql.Open("sqlserver", url.String())
	if err != nil {
		err = fmt.Errorf("while opening database: %v", err)
	}

	return db, err
}
