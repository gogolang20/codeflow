package main

import (
	"database/sql"
	"fmt"
	"net/url"
)

const (
	username = ""
	password = ""
	hostname = ""
	port     = 9000
)

func main() {
	query := url.Values{}
	query.Add("app name", "MyAppName")

	u := &url.URL{
		Scheme: "sqlserver",
		User:   url.UserPassword(username, password),
		Host:   fmt.Sprintf("%s:%d", hostname, port),
		// Path:  instance, // if connecting to an instance instead of a port
		RawQuery: query.Encode(),
	}
	db, err := sql.Open("sqlserver", u.String())
	if err != nil {
		return
	}

	db.Ping()
}
