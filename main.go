package main

import (
	"log"
	"net/http"

	"github.com/ViniciusMartinsS/manager/internal"
	"github.com/ViniciusMartinsS/manager/internal/infrastructure"
	"github.com/ViniciusMartinsS/manager/internal/infrastructure/api"
	"github.com/ViniciusMartinsS/manager/internal/infrastructure/database"
)

var PORT = infrastructure.GetConfig("port")

func main() {
	database.InitDB()
	internal.InitializeDIContainers()

	log.Printf("The server is running on port%s", PORT)
	router := api.SetRoutes()
	log.Fatal(http.ListenAndServe(PORT, router))
}
