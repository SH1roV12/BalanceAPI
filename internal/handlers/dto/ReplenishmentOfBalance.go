package dto

type ReplenishmentOfBalance struct{
    UserID int `json:"user_id" binding:"required,min=1"`
    Amount float64 `json:"amount" binding:"required,gt=0"`
}