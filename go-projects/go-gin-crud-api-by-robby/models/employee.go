package models

import "time"

type Employee struct {
	EmpId          uint `gorm:"primaryKey"`
	EmpName        string
	EmpSalary      float64
	EmpJoiningDate time.Time
}
