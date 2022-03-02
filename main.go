package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	http.HandleFunc("/", anaSayfa)
	http.ListenAndServe(":"+port, nil)
}

func anaSayfa(w http.ResponseWriter, r *http.Request) {
	port := os.Getenv("PORT")
	fmt.Fprintf(w, "Merhaba Dünya!\nKullanılan Port: "+port)
}