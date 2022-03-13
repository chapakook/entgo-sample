package main

import (
	"context"
	"fmt"
	"log"
	"sqlite3/ent"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	client, err := ent.Open("sqlite3", "./pandabank.db?_fk=1")
	checkErr(err)
	defer client.Close()

	ctx := context.Background()
	err = client.Schema.Create(ctx)
	checkErr(err)

	user := client.User.Create().SetID(1).SetName("bob").SaveX(ctx)

	fmt.Println(user)
}

func checkErr(err error) {
	if err != err {
		log.Fatalln(err)
	}
}
