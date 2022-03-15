package database

import (
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"main/block"
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
	var dataStr string

	jsonResp, _ := json.MarshalIndent(peopleData, "", "\n")
	dataStr = string(jsonResp)

	bytes := make([]byte, 32) //AES-256 için rastgele bir 32 bayt anahtar oluşturun.
	if _, err := rand.Read(bytes); err != nil {
		panic(err.Error())
	}
	fmt.Printf("\n")
	key := hex.EncodeToString(bytes) //anahtarı bayt cinsinden kodlayın ve gizli olarak saklayın, bir kasaya koyun
	fmt.Printf("key(anahtar) : %s\n", key)

	encrypted := block.Encrypt(dataStr, key)
	fmt.Printf("encrypted(şifreli) : %s\n", encrypted)

	decrypted := block.Decrypt(encrypted, key)
	fmt.Printf("decrypted(şifre çözüm) : %s\n", decrypted)

	mySlice := []byte(decrypted)

	defer db.Close()
	defer vt.Close()
	w.Write(mySlice)

}
