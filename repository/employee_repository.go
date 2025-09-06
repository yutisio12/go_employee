package repository

import (
	"fmt"
	"gorm.io/gorm"
	"github.com/yutisio12/crud-go-jwt-aes/model"
)

type EmployeeRepository struct { DB *gorm.DB }
func NewEmployeeRepository(db *gorm.DB) *EmployeeRepository {
	return &EmployeeRepository{DB: db}
}

func (r *EmployeeRepository) CreateEmployee(p *model.Employee) error {
	return r.DB.Create(p).Error
}

func (r *EmployeeRepository) FindByBadge(badge string) (*model.Employee, error) {
	var m model.Employee
	if err := r.DB.Where("badge = ?", badge).First(&m).Error; err != nil {
		return nil, err
	}
}