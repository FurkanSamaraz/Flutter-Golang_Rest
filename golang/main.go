package main

import (
	"main/database"
	"net/http"
)

func main() {

	http.HandleFunc("/root", rootPage)

	http.ListenAndServe(":3000", nil)
}

func rootPage(w http.ResponseWriter, r *http.Request) {

	database.GetJson(w, r)
	

}
