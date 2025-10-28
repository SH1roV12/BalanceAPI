package dto

type ReserveUserBalance struct {
	UserId    int     `json:"user_id" binding:"required,min=1"`
	Amount    float64 `json:"amount" binding:"required,min=1"`
	OrderId   int     `json:"order_id"`
	ServiceId int     `json:"service_id"`
}
