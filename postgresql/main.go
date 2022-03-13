package main

import (
	"context"
	"fmt"
	"log"
	"postgresql/ent"

	_ "github.com/lib/pq"
)

func main() {
	client, err := ent.Open("postgres", "host=<host> port=<port> user=<user> dbname=<database> password=<pass>")
	checkErr(err)
	defer client.Close()

	ctx := context.Background()
	err = client.Schema.Create(ctx)
	checkErr(err)

	user := client.User.Create().SetID(1).SetName("bob").SaveX(ctx)

	fmt.Println(user)
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
