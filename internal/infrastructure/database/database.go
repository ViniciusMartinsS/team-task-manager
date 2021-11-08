package database

import (
	"fmt"

	"github.com/ViniciusMartinsS/manager/internal/infrastructure"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	conn             *gorm.DB
	connectionString = infrastructure.GetConfig("database.connection_string")
)

func InitDB() error {
	var err error

	conn, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		fmt.Println("error", err.Error())
		return err
	}

	useCommand := fmt.Sprintf("USE `%s`;", database)
	conn.Exec(useCommand)
	return nil
}

func Connection() *gorm.DB {
	return conn
}
