package handlers

import (
	"errors"
	"log"
	"net/http"
	"users-balance/internal/errmsg"
	"users-balance/internal/handlers/dto"
	"users-balance/internal/models"
	"users-balance/internal/service"

	"github.com/gin-gonic/gin"
)

type UserHandler interface{
	CreateUser(c *gin.Context)
	GetUserBalance(c *gin.Context)
	ReserveUserBalance(c *gin.Context)
	ReplenishmentOfBalance(c *gin.Context) 
}

type Handler struct {
	Service service.UsersService
}


//func for DI
func NewHandler(s service.UsersService) *Handler {
	return &Handler{Service: s}
}


//Create new user method
func (h *Handler) CreateUser(c *gin.Context) {
	var newUser *models.User

	err := c.ShouldBindJSON(&newUser)
	
	if err != nil {
		log.Println("wrong data request")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.Service.CreateUser(c, newUser)


	if err != nil {
		log.Println("Service error",err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	log.Println("successfully creating new user")
	c.JSON(http.StatusCreated, gin.H{"message": "user created successfully"})
}


//Method to getting user balance by id
func (h *Handler) GetUserBalance(c *gin.Context) {
	var req dto.GetUserBalance

	err := c.ShouldBindJSON(&req)


	if err != nil {
		log.Println("wrong data request")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}


	balance, reserved, err := h.Service.GetUserBalance(c.Request.Context(), req)


	if err != nil {

		if errors.Is(err,errmsg.ErrUserNotFound){
			log.Println("user not found by id")
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		log.Println("Service error:", err.Error())

		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})

		return
	}

	log.Println("User data is loaded!")
	c.JSON(http.StatusOK, gin.H{"balance": balance, "reserved": reserved})
}

//Method for reserve user balance
func (h *Handler) ReserveUserBalance(c *gin.Context) {
	var req dto.ReserveUserBalance

	err := c.ShouldBindJSON(&req)


	if err != nil {
		log.Println("wrong data request")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}


	reserved, balance, err := h.Service.ReserveUserBalance(c.Request.Context(), req)


	if err != nil {
		
		if errors.Is(err,errmsg.ErrUserNotFound){
			log.Println("user not found by id")
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		} else if errors.Is(err,errmsg.ErrNotEnoughMoney) {
			log.Println("not enough money on user balance for reservation")
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
			return
		}


		log.Println("Service error:", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	log.Println("Successfully reserved!")
	c.JSON(http.StatusOK, gin.H{"reserved": reserved, "balance": balance})

}


//Method to replenish user balance
func (h *Handler) ReplenishmentOfBalance(c *gin.Context) {
	var req dto.ReplenishmentOfBalance

	err := c.ShouldBindJSON(&req)


	if err != nil {
		log.Println("wrong data request")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}


	balance, err := h.Service.ReplenishmentOfBalance(c.Request.Context(), req)


	if err != nil {

		if errors.Is(err,errmsg.ErrUserNotFound) {
			log.Println("user not found by id")
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		log.Println("Service error:", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
	}

	log.Println("Successfully balance replenished!")
	c.JSON(http.StatusOK, gin.H{"enrolled": req.Amount, "balance": balance})
}
