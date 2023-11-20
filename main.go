package main

import (
	"fmt"
	"net/http"

	"./common"
	_ "github.com/lib/pq"
)

func homePage(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "Connected to the database")
	fmt.Fprintf(w, "Value %d", common.QueryDB())
}

func setupRoutes() {
	http.HandleFunc("/", homePage)
}

func main() {
	fmt.Println("Go Web App Started on Port 3000")
	setupRoutes()
	http.ListenAndServe(":3000", nil)
}
