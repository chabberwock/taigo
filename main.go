package main

import (
	"net/http"
	"log"
	"database/sql"
)

var db *sql.DB

func main() {

	db, _ = sql.Open("mysql", "root:@/passport")
//	importProgress := make(chan string)
//	downloadProgress := make(chan string)

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/parse", parseHandler)
	http.HandleFunc("/check", checkHandler)


	log.Fatal(http.ListenAndServe(":8080", nil))

}
