package config

import "os"

var Token = os.Getenv("TOKEN")
var Address = os.Getenv("RETHINKDB_HOST")
var Database = os.Getenv("RETHINKDB_NAME")
var Username = os.Getenv("RETHINKDB_USERNAME")
var Password = os.Getenv("RETHINKDB_PASSWORD")
