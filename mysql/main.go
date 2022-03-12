package main

import (
	"log"

	"entgo.io/ent/entc/integration/edgefield/ent"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	client, oErr := ent.Open("mysql", "admin:admin@tcp(127.0.0.1:3306)/pandabank?parseTime=True")
	checkErr(oErr)
	defer client.Close()
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
