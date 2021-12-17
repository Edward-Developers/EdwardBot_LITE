package database

import (
	"fmt"
	"EdwardBot_LITE/structs"
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

func CreateTables(session *r.Session) {
	var tableNames = []structs.Table{
		{ NAME: "Settings", MULTI: true, INDEXES: []structs.Indexes{
			{ NAME: "ID" },
		} },
		{ NAME: "Economies", MULTI: false },
	}
    for _, table := range tableNames {
		if !table.MULTI {
			r.TableCreate(table.NAME).Run(session)
			fmt.Println("Create the table: "+ table.NAME)
		} else {
			for _, index := range table.INDEXES {
				r.TableCreate(table.NAME).Run(session)
				r.Table(table.NAME).IndexCreate(index.NAME).Run(session)
				r.Table(table.NAME).IndexWait()
			}
			fmt.Println("Create the table: "+ table.NAME)
		}
    }
}