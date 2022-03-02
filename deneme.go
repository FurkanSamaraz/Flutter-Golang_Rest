package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("pages")))
	http.HandleFunc("/players", getPlayers)
	log.Println("Web server is active at port 8045")
	http.ListenAndServe(":8045", nil)
}

func getPlayers(response http.ResponseWriter, request *http.Request) {
	log.Println("Get Request for Players")
	players := loadPlayers()
	log.Println("Players loaded")
	t, err := template.ParseFiles("pages/Players.html")
	if err == nil {
		t.Execute(response, players)
	} else {
		log.Println(err.Error())
	}
}

type Player struct {
	Id       int
	Kullancı string
	Level    int
}

func loadPlayers() []Player {
	return []Player{
		Player{1001, "Eren", 10},
		Player{1002, "Furkan", 6},
	}
}
