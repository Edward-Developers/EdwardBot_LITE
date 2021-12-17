package structs

type Settings struct {
	ID string `rethinkdb:"id"`
	PREFIX string
}