# golearn - ent (mac M1)


### go module
go module setting

GO111MODULE state : off, on, auto(default)
- off  : build package use $GOPATH
- on   : build package use project module(go.mod)
- auto : build package if project is a module(go.mod) use it, there is no module(go.mod) use $GOPATH
```
go env -w GO111MODULE=[state]
```
-----

### go get option
link : https://pkg.go.dev/cmd/go/internal/get

```
go get [-d] [-f] [-t] [-u] [-v] [-fix] [-insecure] [build flags] [packages]
```
-----
### ent doc
link : https://entgo.io/docs/getting-started/
 
-----
### go module create -> 'go.mod'
```
go mod init <project>
```

### go module ent add -> 'go.mod'
```
go get entgo.io/ent/cmd/ent
go get -d entgo.io/ent/cmd/ent
```
-----
### ent schema init
```
go run entgo.io/ent/cmd/ent init <schema> <schema> <>...
```
-----
### schema setting
link : https://entgo.io/docs/schema-def

-----
### generate schema
```
go generate ./ent
```
-----
### mysql
1. driver add
```
go get github.com/go-sql-driver/mysql
go get -d github.com/go-sql-driver/mysql
```
2. open
```
import (
    ...
    _ "github.com/go-sql-driver/mysql"
)

func main() {
    client, err := ent.Open("mysql", "<user>:<password>@<protocol>(<host>:<port>)/<database>?parseTime=True")
    if err != nil {
        log.Fatal(err)
    }
    defer client.Close()
}
```
-----
### SQLite3
1. driver add
```
go get github.com/mattn/go-sqlite3
go get -d github.com/mattn/go-sqlite3
```

2. open
```
import (
    ...
    _ "github.com/mattn/go-sqlite3"
)

func main() {
    client, err := ent.Open("sqlite3", "file:<file.db>?_fk=1")
    if err != nil {
        log.Fatal(err)
    }
    defer client.Close()
}
```
-----
### ~~PostgreSQL~~ (here not use)
1. driver add
```
go get github.com/lib/pq
go get -d github.com/lib/pq
```

2. open
```
import (
    ...
    _ "github.com/lib/pq"
)

func main() {
    client, err := ent.Open("postgres","host=<host> port=<port> user=<user> dbname=<database> password=<password>")
    if err != nil {
        log.Fatal(err)
    }
    defer client.Close()
}
```
-----
