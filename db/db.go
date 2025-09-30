package db

import (
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
	
)



func ConnectionDB() *sql.DB {
	db, err := sql.Open("mysql", "root:yangi_parol@tcp(127.0.0.1:3306)/Auth")
	if err != nil {
		log.Fatal("database connection in error", err)
	}

   err = db.Ping()
   if err != nil {
	log.Fatal("database ping query error", err)
   }
   return  db


}