package database

import (
	"log"
	"fmt"
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

func ConnectTest(a int, b int) int {
	return a + b
}
func Connect() {
	session, err := r.Connect(r.ConnectOpts{
		Address: "localhost:28015",
		Database: "test",
		AuthKey:  "14daak1cad13dj",
	})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(session)
}