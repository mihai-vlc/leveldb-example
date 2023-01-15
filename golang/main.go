package main

import (
	"encoding/json"
	"log"

	"github.com/syndtr/goleveldb/leveldb"
)

type User struct {
	Name string `json:"name"`
}

type Data struct {
	Account User `json:"user"`
}

func main() {
	db, err := leveldb.OpenFile("../mydata", nil)

	if err != nil {
		log.Fatalf("Failed to open db: %s", err)
	}

	data, err := db.Get([]byte("name"), nil)

	if err != nil {
		log.Fatalf("Failed to read value: %s", err)
	}

	log.Printf("data: %s \n", data)

	var d1 = Data{}

	json.Unmarshal(data, &d1)

	log.Printf("d1: %s", d1)
	log.Printf("account.name: %s", d1.Account.Name)
}
