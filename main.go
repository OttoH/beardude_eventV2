package main

import (
	"github.com/gin-gonic/gin"

	"beardude_eventV2/config"
	"beardude_eventV2/dao"
)

var cfg = &config.Config{}

func main() {
	// read configuaration
	cfg.Read()

	// setup Mongo DB
	db := &dao.DAO{
		Server:   cfg.Server,
		Database: cfg.Database,
	}
	session := db.Connect()
	defer session.Close()

	// Disable Console Color
	// gin.DisableConsoleColor()
	// r := gin.New() // without middleware

	authMiddleware := GetAuthMiddleware(db)

	r := gin.Default()

	server := SetupRouter(r, db, authMiddleware)
	server.Run(":8080")
}
