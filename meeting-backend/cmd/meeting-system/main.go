package main

import (
	"backend/internal/database"
	"backend/internal/server"
)

func main() {
	database.DB()
	server.New().Run()
}
