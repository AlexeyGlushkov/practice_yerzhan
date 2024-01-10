package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Creates an employee and position
// @Description Method for creating a new employee and his position
// @Tags Employees
// @Accept json
// @Produce json
// @Param body body CreateEmployeePayload true "New details"
// @Success 201 {string} string "Employee and Position created successfully"
// @Failure 400 {string} string "Invalid request payload"
// @Failure 500 {string} string "Internal server error"
// @Router /employee [post]
func CreateEmployeeHandler(svc *Service) gin.HandlerFunc {
	return func(c *gin.Context) {

		var payload CreateEmployeePayload

		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := svc.CreateService(c.Request.Context(), payload.Employee, payload.Position); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "Employee and Position created successfully"})
	}
}

// @Summary Return Employee details by employeeID
// @Description Method for retrieving employee details by ID
// @Tags Employees
// @Accept json
// @Produce json
// @Param id path string true "EmployeeID"
// @Success 200 {object} Employee "Employee details"
// @Failure 500 {string} string "Internal server error"
// @Failure 404 {string} string "Employee not found"
// @Router /employee/{id} [get]
func GetEmployeeHandler(svc *Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		employeeID := c.Param("id")

		employee, err := svc.GetByIDService(c.Request.Context(), employeeID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if employee.Employee_id == "" {
			c.JSON(http.StatusNotFound, gin.H{"message": "Employee not found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"employee": employee})
	}
}

// @Summary Deletes Employee details by employeeID
// @Description Method for deleting employee details by ID
// @Tags Employees
// @Accept json
// @Produce json
// @Param id path string true "EmployeeID"
// @Success 200 {string} string "Employee deleted successfully"
// @Failure 500 {string} string "Internal server error"
// @Router /employee/{id} [delete]
func DeleteEmployeeHandler(svc *Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		employeeID := c.Param("id")

		err := svc.DeleteService(c.Request.Context(), employeeID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Employee deleted successfully"})
	}
}

// @Summary Updates Employee details
// @Description Method for updating employee details
// @Tags Employees
// @Accept json
// @Produce json
// @Param id path string true "EmployeeID"
// @Param body body UpdateEmployeePayload true "New details"
// @Success 200 {string} string "Employee updated successfully"
// @Failure 500 {string} string "Internal server error"
// @Failure 400 {string} string "Invalid request payload"
// @Router /employee/{id} [put]
func UpdateEmployeeHandler(svc *Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		employeeID := c.Param("id")

		var updateData UpdateEmployeePayload

		if err := c.ShouldBindJSON(&updateData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		err := svc.UpdateEmployeeService(c.Request.Context(), employeeID, updateData.FirstName, updateData.LastName)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Employee updated successfully"})
	}
}
