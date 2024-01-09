package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateEmployeeHandler(svc *Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			Employee Employee `json:"employee"`
			Position Position `json:"position"`
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := svc.CreateService(c.Request.Context(), req.Employee, req.Position); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "Employee and Position created successfully"})
	}
}

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

func UpdateEmployeeHandler(svc *Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		employeeID := c.Param("id")

		var updateData struct {
			FirstName string `json:"first_name"`
			LastName  string `json:"last_name"`
		}

		if err := c.ShouldBindJSON(&updateData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		employee, err := svc.UpdateEmployeeService(c.Request.Context(), employeeID, updateData.FirstName, updateData.LastName)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"employee": employee})
	}
}
