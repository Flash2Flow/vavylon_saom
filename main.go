package main

import (
	"log"
	"net/http"
	"os"
)


func main() {
	//design pages
	http.HandleFunc("/", HomePage)

	//tech pages

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	port := os.Getenv("PORT")
	log.Println("start server with port: " +port)
	http.ListenAndServe(":"+port, nil)

}
