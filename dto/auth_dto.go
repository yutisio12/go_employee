package dto

type LoginRequest struct {
	BadgeId uint   `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}