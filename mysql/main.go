package main

import (
	"context"
	"fmt"
	"log"
	"mysql/ent"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	client, err := ent.Open("mysql", "admin:admin@tcp(127.0.0.1:3306)/pandabank?parseTime=True")
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
