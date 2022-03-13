#!/bin/sh

echo echo "Start init postgresql entgo"
go mod init postgresql

echo "GET entgo.io/ent/cmd/ent"
go get -d entgo.io/ent/cmd/ent

echo "GET github.com/lib/pq"
go get -d github.com/lib/pq

echo "Init Schema"
go run entgo.io/ent/cmd/ent init User Account

echo "cp Schema 'User.go' 'Account.go'"
cp -r schema ent

echo "go generate"
go generate ./ent