package database

import (
	"fmt"
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

func CreateTables(session *r.Session) {
	var tableNames = [2]string{"Settings", "Economies"}
    for _, name := range tableNames {
        r.TableCreate(name).Run(session)
		fmt.Println("Create the table: "+ name)
    }
}