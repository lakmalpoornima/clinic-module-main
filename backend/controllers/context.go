package controllers

import (
	"context"
	"database/sql"
	"fmt"
	"net/url"
	"os"
	"time"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

var ctx context.Context
var LogFile *os.File
var db *sql.DB

func init() {
	query := url.Values{}
	query.Add("database", os.Getenv("database_name"))
	dns := &url.URL{
		Scheme:   "sqlserver",
		User:     url.UserPassword(os.Getenv("database_user"), os.Getenv("database_password")),
		Host:     fmt.Sprintf("%s:%s", os.Getenv("database_ip"), os.Getenv("database_port")),
		RawQuery: query.Encode(),
	}
	mssqlconfig := dns.String()
	ctx = context.Background()
	var err error
	db, err = sql.Open("sqlserver", mssqlconfig)
	if err != nil {
		print(err.Error())
	}
	boil.DebugMode = true
	LogFile, err = os.OpenFile("logs.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	boil.DebugWriter = LogFile
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	db.SetMaxIdleConns(25)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	db.SetMaxOpenConns(25)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	db.SetConnMaxLifetime(5 * time.Minute)
}
