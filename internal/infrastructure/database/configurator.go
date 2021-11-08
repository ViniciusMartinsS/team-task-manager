package database

import (
	"fmt"
	"log"

	"github.com/ViniciusMartinsS/manager/internal/domain"
	"github.com/ViniciusMartinsS/manager/internal/domain/contract"
	"github.com/ViniciusMartinsS/manager/internal/infrastructure"
	"gorm.io/gorm"
)

type connection struct {
	db *gorm.DB
}

var (
	database  = infrastructure.GetConfig("database.name")
	roleTable = infrastructure.GetConfig("database.table.roles")
	taskTable = infrastructure.GetConfig("database.table.tasks")
	userTable = infrastructure.GetConfig("database.table.users")
)

func NewDatabaseConfigurator(db *gorm.DB) contract.Database {
	return &connection{db}
}

func (c connection) CreateDatabase() error {
	createCommand := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s`;", database)
	c.db = c.db.Exec(createCommand)
	if c.db.Error != nil {
		return c.db.Error
	}

	useCommand := fmt.Sprintf("USE `%s`;", database)
	c.db = c.db.Exec(useCommand)
	if c.db.Error != nil {
		return c.db.Error
	}

	log.Println("Database create successfully!")

	return nil
}

func (c connection) CreateTables() error {
	var err error

	if err = c.db.Migrator().DropTable(roleTable); err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	if err = c.db.Migrator().CreateTable(&domain.Role{}); err != nil {
		return c.db.Error
	}

	if err = c.db.Migrator().DropTable(taskTable); err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	if err = c.db.Migrator().CreateTable(&domain.Task{}); err != nil {
		return c.db.Error
	}

	if err = c.db.Migrator().DropTable(userTable); err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	if err = c.db.Migrator().CreateTable(&domain.User{}); err != nil {
		return c.db.Error
	}

	log.Println("Tables created successfully!")

	return nil
}

func (c connection) SeedTables() {
	for _, role := range roles {
		c.db.Model(&domain.Role{}).Create(&role)
	}

	for _, task := range tasks {
		c.db.Model(&domain.Task{}).Create(&task)
	}

	for _, user := range users {
		c.db.Model(&domain.User{}).Create(&user)
	}
}
