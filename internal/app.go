package internal

import (
	"users-balance/internal/database"
	"users-balance/internal/handlers"
	"users-balance/internal/repository"
	"users-balance/internal/service"

	"github.com/gin-gonic/gin"
)



func StartApp(){
	r:=gin.Default()
	db:=database.InitDB()
	repo:=repository.NewRepository(db)
	srvc:=service.NewService(repo)
	handler:=handlers.NewHandler(srvc)
	

	r.POST("/users/create",handler.CreateUser)
	r.PUT("/users/reserve",handler.ReserveUserBalance)
	r.GET("/balance", handler.GetUserBalance)
	r.PUT("/balance",handler.ReplenishmentOfBalance)
	r.Run(":8080")
}