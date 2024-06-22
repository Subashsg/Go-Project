package handlers

import (
	"employeeattenance/dbconnection"
	"employeeattenance/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CheckIn(c *gin.Context) {
	var employee models.Employee
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	now := time.Now()
	employee.CheckIn = now
	employee.CheckOut = nil
	employee.HoursWorked = 0
	if err := dbconnection.GetDB().Create(&employee).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, employee)
}

func CheckOut(c *gin.Context) {
	var employee models.Employee
	id := c.Param("id")
	if err := dbconnection.GetDB().Where("id = ?", id).First(&employee).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
		return
	}
	now := time.Now()
	employee.CheckOut = &now
	duration := now.Sub(employee.CheckIn)
	employee.HoursWorked = int(duration.Seconds())
	if err := dbconnection.GetDB().Save(&employee).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, employee)
}

func AttendanceReport(c *gin.Context) {
	var employee models.Employee
	id := c.Param("id")
	if err := dbconnection.GetDB().Where("id = ?", id).First(&employee).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
		return
	}
	c.JSON(http.StatusOK, employee)
}

func FilteredAttendance(c *gin.Context) {
	var employees []models.Employee
	timeFilter := c.Query("time")
	hoursFilter := c.Query("hours")

	if timeFilter != "" && hoursFilter != "" {
		parsedTime, err := time.Parse(time.RFC3339, timeFilter)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid time format"})
			return
		}
		parsedHours, err := time.ParseDuration(hoursFilter + "h")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid hours format"})
			return
		}
		dbconnection.GetDB().Where("check_in >= ? AND hours_worked >= ?", parsedTime, parsedHours.Hours()).Find(&employees)
	} else {
		dbconnection.GetDB().Find(&employees)
	}

	c.JSON(http.StatusOK, employees)
}
