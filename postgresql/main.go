package main

import (
	"log"

	"entgo.io/ent/entc/integration/edgefield/ent"
	_ "github.com/lib/pq"
)

func main() {
	client, err := ent.Open("postgres", "host=<host> port=<port> user=<user> dbname=<database> password=<pass>")
	checkErr(err)
	defer client.Close()
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
