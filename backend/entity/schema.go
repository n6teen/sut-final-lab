package entity

import (
	"gorm.io/gorm"
)

type Employees struct {
	gorm.Model
	Name       string
	URL        string `valid:"url~URL is invalid"`
	EmployeeID string `valid:"matches(^EM\\d{10}$)~EmployeeID is invalid"`
}
