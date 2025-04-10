package main

import (
	"log"
	"server/config"
	"server/db"
	"server/internal/user"
	"server/internal/ws"
	"server/router"
)

func main() {
	config.LoadConfig()

	dbConn, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("Error creating database: %s", err)
	}

	userRep := user.NewRepository(dbConn.GetDB())
	userSvc := user.NewService(userRep)
	userHandler := user.NewHandler(userSvc)

	hub := ws.NewHub()
	wsHandler := ws.NewHandler(hub)
	go hub.Run()

	router.InitRouter(userHandler, wsHandler)
	router.Start("0.0.0.0:8080")
}
