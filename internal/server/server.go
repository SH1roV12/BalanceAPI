package server

import (
	"log"
	"users-balance/internal/handlers"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Router  *gin.Engine
	Handler *handlers.Handler
}

//func for DI
func NewServer(h *handlers.Handler) *Server {
	return &Server{
		Router:  gin.Default(),
		Handler: h,
	}
}

//Registration routes
func (s *Server) Routes() {

	s.Router.POST("/users/create", s.Handler.CreateUser)

	s.Router.PUT("/users/reserve", s.Handler.ReserveUserBalance)

	s.Router.GET("/balance", s.Handler.GetUserBalance)

	s.Router.PUT("/balance", s.Handler.ReplenishmentOfBalance)
}

//Starting server
func (s *Server) RunServer(port string){
	log.Println("Starting httpserver...")
	s.Routes()
	err := s.Router.Run(port)
	if err != nil {
		log.Fatalln("Failed to starting server")
	}
	log.Println("Http server is running")
	
}
