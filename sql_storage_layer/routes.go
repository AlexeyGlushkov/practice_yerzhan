package main

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter(svc *Service) *gin.Engine {

	router := gin.Default()

	employeeGroup := router.Group("/v1/employee")
	{
		employeeGroup.POST("/", CreateEmployeeHandler(svc)) // Добавить возвращение employeeID в месседже
		employeeGroup.GET("/:id", GetEmployeeHandler(svc))
		employeeGroup.PUT("/:id", UpdateEmployeeHandler(svc)) // Status 200, но records отсутствуют
		employeeGroup.DELETE("/:id", DeleteEmployeeHandler(svc))
	}

	return router
}
