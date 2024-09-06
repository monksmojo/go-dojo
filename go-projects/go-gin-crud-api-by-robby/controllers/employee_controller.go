package controllers

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/monksmojo/go-dojo/go-projects/go-gin-crud-api-by-robby/dto"
	"github.com/monksmojo/go-dojo/go-projects/go-gin-crud-api-by-robby/initializers"
	"github.com/monksmojo/go-dojo/go-projects/go-gin-crud-api-by-robby/models"
	"gorm.io/gorm"
)

func CreateEmployee(ginContext *gin.Context) {
	// bind data from request body
	var employeeDTO dto.EmployeeDTO
	ginContext.Bind(&employeeDTO)
	// create new employee
	newEmployeeId := getLastEmployeeId() + 1
	newEmployee := models.Employee{
		EmpId:          newEmployeeId,
		EmpName:        employeeDTO.EmpName,
		EmpSalary:      employeeDTO.EmpSalary,
		EmpJoiningDate: time.Now(),
	}
	result := initializers.DB.Create(newEmployee)
	if result.Error != nil {
		msg := fmt.Sprintf("internal server error %v \n", result.Error)
		ginContext.JSON(
			500, gin.H{
				"message": msg,
			})
	}
	msg := fmt.Sprintf("employee created with id %v \n", newEmployeeId)
	ginContext.JSON(200, gin.H{
		"response": msg,
	})
}

func getLastEmployeeId() uint {
	var employee models.Employee
	result := initializers.DB.Last(&employee)
	if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return 1001
	} else {
		log.Fatal("error fetching the last record", result.Error)
	}
	return employee.EmpId
}
