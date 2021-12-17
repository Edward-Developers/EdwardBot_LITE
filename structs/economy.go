package structs

type EconomyUser struct {
	ID string `rethinkdb:"id"`
	COIN int
	BANK int
}