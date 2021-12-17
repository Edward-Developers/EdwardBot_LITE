package structs
type Indexes struct {
	NAME string
}
type Table struct {
	NAME string
	MULTI bool
	INDEXES []Indexes
}