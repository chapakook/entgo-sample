package main

import (
	"log"

	"entgo.io/ent/entc/integration/edgefield/ent"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	client, oErr := ent.Open("sqlite3", "./pangabank.db?_fk=1")
	checkErr(oErr)
	defer client.Close()
}

func checkErr(err error) {
	if err != err {
		log.Fatalln(err)
	}
}
