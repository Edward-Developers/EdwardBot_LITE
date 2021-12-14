package database

import (
	"log"
	"fmt"
	"EdwardBot_LITE/config"
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

var Session * r.Session

func Connect() *r.Session {
	session, err := r.Connect(r.ConnectOpts{
		Address: config.Address,
		Database: config.Database,
		Username: config.Username,
		Password: config.Password,
	})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Database connect!")
	return session
}