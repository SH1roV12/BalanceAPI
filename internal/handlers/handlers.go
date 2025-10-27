package handlers

import (
	"log"
	"net/http"
	"users-balance/internal/handlers/dto"
	"users-balance/internal/models"
	"users-balance/internal/service"

	"github.com/gin-gonic/gin"
)


type Handler struct{
	Service *service.UsersService
}

func NewHandler(s *service.UsersService)*Handler{
	return &Handler{Service: s}
}


func (h *Handler) CreateUser(c *gin.Context){
	var newUser *models.User
	err := c.ShouldBindJSON(&newUser)
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}
	err=h.Service.CreateUser(c,newUser)
	if err !=nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		return
	}
	c.JSON(http.StatusCreated,gin.H{"message":"user created successfully"})
}


func (h *Handler) GetUserBalance(c *gin.Context){
	var req dto.GetUserBalance
	err:=c.ShouldBindJSON(&req)
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	balance,reserved,err:=h.Service.GetUserBalance(c.Request.Context(),req)
	if err!=nil{
			if err.Error() == "user not found"{
			c.JSON(http.StatusNotFound,gin.H{"error":err.Error()})
			return
		}
		log.Println("Service error:",err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error":"internal server error"})
		return
	}
	c.JSON(http.StatusOK,gin.H{"balance":balance, "reserved":reserved})
}

func (h *Handler) ReserveUserBalance(c *gin.Context){
	var req dto.ReserveUserBalance
	err:=c.ShouldBindJSON(&req)
	if err !=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}
	balance,reserved,err:=h.Service.ReserveUserBalance(c.Request.Context(),req)
	if err!=nil{
		if err.Error() == "user not found" {
			c.JSON(http.StatusNotFound,gin.H{"error":err.Error()})
			return
		}else if err.Error() == "not enough money" {
			c.JSON(http.StatusForbidden,gin.H{"error":err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error":"internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"reserved":reserved, "balance":balance})

}


func (h *Handler) ReplenishmentOfBalance(c *gin.Context){
	var req dto.ReplenishmentOfBalance
	err:=c.ShouldBindJSON(&req)
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}
	balance,err:=h.Service.ReplenishmentOfBalance(c.Request.Context(),req)
	if err!= nil{
			if err.Error() == "user not found"{
			c.JSON(http.StatusNotFound,gin.H{"error":err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError,gin.H{"error":"internal server error"})
	}
	
	c.JSON(http.StatusOK,gin.H{"enrolled":req.Amount,"balance":balance})
}