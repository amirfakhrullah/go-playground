package models

import (
	"time"
	"github.com/amirfakhrullah/go-playground/tree/main/hr-management/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Employee struct {
	gorm.Model
	FullName    string    `json:"fullName"`
	Email       string    `gorm:"index:unique" json:"email"`
	Designation string    `json:"designation"`
	Salary      float64   `json:"salary"`
	AnnualLeave int16     `json:"annualLeave"`
	Leaves      Leave     `gorm:"references:ID"`
	StartDate   time.Time `json:"startDate"`
	EndDate     time.Time `json:"endTime"`
}

type Leave struct {
	gorm.Model
	StartDate time.Time `json:"startDate"`
	EndTime   time.Time `json:"endTime"`
}

func init() {
	config.Connect()
	db = config.GetDb()
	db.AutoMigrate(&Employee{}, &Leave{})
}

func (employee *Employee) CreateEmployee() *Employee {
	db.Create(&employee)
	return employee
}

func GetAllEmployees() []Employee {
	var employees []Employee
	db.Find(&employees)
	return employees
}

func GetEmployeeById(id int64) (*Employee, *gorm.DB) {
	var employee Employee
	db := db.Where("ID=?", id).Find(&employee)
	return &employee, db
}

func DeleteEmployee(id int64) Employee {
	var employee Employee
	db.Where("ID=?", id).Delete(&employee)
	return employee
}
