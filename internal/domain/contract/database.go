package contract

type Database interface {
	CreateDatabase() error
	CreateTables() error
	SeedTables()
}
