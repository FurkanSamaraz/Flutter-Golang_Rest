package database

import (
	"database/sql"
	"encoding/json"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

type Data struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

var db *sql.DB

func GetJson(w http.ResponseWriter, r *http.Request) {
	db, _ = sql.Open("sqlite3", "veri.db")

	vt, _ := db.Query("SELECT * FROM data")
	var dataUs Data
	var peopleData []Data

	for vt.Next() {

		vt.Scan(&dataUs.Id, &dataUs.Name)

		peopleData = append(peopleData, dataUs)

	}
	jsonResp, _ := json.MarshalIndent(peopleData, "", "\n")
	w.Write(jsonResp)

	defer db.Close()
	defer vt.Close()
	return

}

func GGG(w http.ResponseWriter, r *http.Request) {
	GetJson(w, r)

}
