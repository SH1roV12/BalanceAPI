package server

import (
	"users-balance/internal/handlers"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Router  *gin.Engine
	Handler *handlers.Handler
}

func NewServer(h *handlers.Handler) *Server {
	return &Server{
		Router:  gin.Default(),
		Handler: h,
	}
}

func (s *Server) Routes() {
	s.Router.POST("/users/create", s.Handler.CreateUser)
	s.Router.PUT("/users/reserve", s.Handler.ReserveUserBalance)
	s.Router.GET("/balance", s.Handler.GetUserBalance)
	s.Router.PUT("/balance", s.Handler.ReplenishmentOfBalance)

}

func (s *Server) RunServer(port string) error {
	s.Routes()
	err := s.Router.Run(port)
	if err != nil {
		return err
	}
	return nil
}
