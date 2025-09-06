package dto

type CreateEmployee struct {
	Name  string `json:"name" binding:"required"`
	Badge int    `json:"badge" binding:"required"`
	Mail  string `json:"mail" binding:"required"`
}

type UpdateEmployee struct {
	Name  string `json:"name"`
	Badge int    `json:"badge"`
	Mail  string `json:"mail"`
}

type EmployeeResponse struct {
	BadgeId uint   `json:"badge_id"`
	Name    string `json:"name"`
	Mail    string `json:"mail"`
}