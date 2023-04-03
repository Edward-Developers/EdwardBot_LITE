package database

import (
	"EdwardBot_LITE/structs"
	"fmt"
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

func CreateTables(session *r.Session) {
	var tableNames = []structs.Table{
		{NAME: "Settings", MULTI: true, INDEXES: []structs.Indexes{
			{NAME: "ID"},
		}},
		{NAME: "Economies", MULTI: false},
	}
	for _, table := range tableNames {
		if !table.MULTI {
			_, err := r.TableCreate(table.NAME).Run(session)
			if err != nil {
				return
			}
			fmt.Println("Create the table: " + table.NAME)
		} else {
			for _, index := range table.INDEXES {
				_, err := r.TableCreate(table.NAME).Run(session)
				if err != nil {
					return
				}
				_, err = r.Table(table.NAME).IndexCreate(index.NAME).Run(session)
				if err != nil {
					return
				}
				r.Table(table.NAME).IndexWait()
			}
			fmt.Println("Create the table: " + table.NAME)
		}
	}
}
