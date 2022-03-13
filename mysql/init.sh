#!/bin/sh

echo echo "Start init mysql entgo"
go mod init mysql

echo "GET entgo.io/ent/cmd/ent"
go get -d entgo.io/ent/cmd/ent

echo "GET github.com/go-sql-driver/mysql"
go get -d github.com/go-sql-driver/mysql

echo "Init Schema"
go run entgo.io/ent/cmd/ent init User Account

echo "cp Schema 'User.go' 'Account.go'"
cp -r schema ent

echo "go generate"
go generate ./ent