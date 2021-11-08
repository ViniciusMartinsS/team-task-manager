package main

import (
	"fmt"

	"github.com/ViniciusMartinsS/manager/internal/infrastructure"
	"github.com/ViniciusMartinsS/manager/internal/infrastructure/database"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	connectionString := infrastructure.GetConfig("database.connection_string")

	conn, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		fmt.Println("error", err.Error())
		panic(err)
	}

	seed := database.NewDatabaseConfigurator(conn)

	err = seed.CreateDatabase()
	if err != nil {
		panic(err)
	}

	err = seed.CreateTables()
	if err != nil {
		panic(err)
	}

	seed.SeedTables()
}
