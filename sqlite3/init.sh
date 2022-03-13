#!/bin/sh

echo "Start init sqlite3 entgo"
go mod init sqlite3

echo "GET entgo.io/ent/cmd/ent"
go get -d entgo.io/ent/cmd/ent

echo "GET github.com/mattn/go-sqlite3"
go get -d github.com/mattn/go-sqlite3

echo "Init Schema"
go run entgo.io/ent/cmd/ent init User Account

echo "cp Schema 'User.go' 'Account.go'"
cp Schema ent/schema

echo "go generate"
go generate ./ent