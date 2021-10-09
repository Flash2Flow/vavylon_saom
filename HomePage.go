package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func HomePage(page http.ResponseWriter, req *http.Request) {
	temp, err := template.ParseFiles("templates/html/HomePage.html")

	if err != nil {
		fmt.Fprintf(page, err.Error())
	}

	temp.ExecuteTemplate(page, "home_page", nil)
}
