package main

import (
	"context"
	"fmt"
	"log"
	"mysql/ent"

	_ "github.com/go-sql-driver/mysql"
)

type owner struct {
	id    int
	name  string
	money int
}

func main() {
	client, err := ent.Open("mysql", "admin:admin@tcp(127.0.0.1:3306)/pandabank?parseTime=True")
	checkErr(err)
	defer client.Close()

	ctx := context.Background()
	err = client.Schema.Create(ctx)
	checkErr(err)

	owner := []owner{
		{id: 1, name: "kaka", money: 100},
		{id: 2, name: "nunu", money: 10},
		{id: 3, name: "santo", money: 200},
		{id: 4, name: "gura", money: 1000},
		{id: 5, name: "yuon", money: 500},
	}

	for _, o := range owner {
		_, err := client.User.Create().SetID(o.id).SetName(o.name).Save(ctx)
		checkErr(err)
		_, err = client.Account.Create().SetID(o.id).SetMoney(o.money).SetOwnerID(o.id).Save(ctx)
		checkErr(err)
	}

	as := client.Account.Query().AllX(ctx)

	for _, a := range as {
		fmt.Println(a)
		err := client.Account.DeleteOneID(a.ID).Exec(ctx)
		checkErr(err)
	}

	us := client.User.Query().AllX(ctx)

	for _, u := range us {
		fmt.Println(u)
		err := client.User.DeleteOneID(u.ID).Exec(ctx)
		checkErr(err)
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
