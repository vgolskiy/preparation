package main

import (
	_ "github.com/lib/pq"
	H "ozon/src/handlers"
	"log"
	M "ozon/src/model"
	"net/http"
)

func main() {
	M.DBConnect()
	defer M.DB.Close()

	http.HandleFunc("/", H.Handler)
	log.Fatal(http.ListenAndServe(":8888", nil))
}
