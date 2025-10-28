package dto

//DTO for GetUserBalance method
type GetUserBalance struct {
	UserId int `json:"user_id" binding:"required,min=1"`
}
