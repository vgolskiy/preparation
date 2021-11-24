package main

import (
	H "./handlers" //H - handlers for request/response
	M "./model"	//M - DB model functions
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func main() {
	M.DBConnect()
	defer M.DB.Close()

	http.HandleFunc("/", H.Handler)
	log.Fatal(http.ListenAndServe(":8888", nil))
}
