package storage

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db   *sql.DB
	once sync.Once
)

/*
!Where We were working with date we must implements "?parseTime=true" on of the script of connection had our database
*/
const (
	dns = "usi1p2fhaqjrbcxd:7b3gioYphTLYaZoRbF8P@tcp(byr2cqsjyxldxj3oabah-mysql.services.clever-cloud.com)/byr2cqsjyxldxj3oabah?parseTime=true"
)

func NewConnection() {
	once.Do(func() {
		var err error
		db, err = sql.Open("mysql", dns)
		if err != nil {
			log.Fatalf("Can't open database %v", err)
		}

		if err = db.Ping(); err != nil {
			log.Fatalf("Can't do ping %v", err)
		}

		fmt.Println("Conecction successfully!!!")
	})
}

func Pool() *sql.DB {
	return db
}

/*
? function for validated the value null
*/
func stringToNull(s string) sql.NullString {
	null := sql.NullString{String: s}
	if null.String != "" {
		null.Valid = true
	}
	return null
}
