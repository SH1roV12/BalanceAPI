package models

//main model of user
type User struct {
	ID       int     `json:"id"`
	Balance  float64 `json:"balance"`
	Reserved float64 `json:"reserved"`
}
