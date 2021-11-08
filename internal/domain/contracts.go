package domain

type Database interface {
	CreateDatabase() error
	CreateTables() error
	SeedTables()
}
