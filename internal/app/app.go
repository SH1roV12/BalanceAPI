package app

import (
	"users-balance/internal/database"
	"users-balance/internal/handlers"
	"users-balance/internal/repository"
	"users-balance/internal/server"
	"users-balance/internal/service"
)

func StartApp() {
	db := database.InitDB()
	repo := repository.NewRepository(db)
	srvc := service.NewService(repo)
	handler := handlers.NewHandler(srvc)
	httpServer := server.NewServer(handler)
	httpServer.RunServer(":8080")
}
