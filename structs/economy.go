package structs

type EconomyUser struct {
	ID string `rethinkdb:"id"`
	COIN int64
	BANK int64
}
type EconomyUserUpdate struct {
	COIN int64
	BANK int64
}