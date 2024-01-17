package http

import (
	servc "sql_storage_layer/pkg/service"

	"github.com/gin-gonic/gin"
)

func SetupRouter(svc *servc.Service) *gin.Engine {

	router := gin.New()

	router.Use(LoggerMiddleware())

	employeeGroup := router.Group("/v1/employee")
	{
		employeeGroup.POST("/", CreateEmployeeHandler(svc)) // Добавить возвращение employeeID в месседже
		employeeGroup.GET("/:id", GetEmployeeHandler(svc))
		employeeGroup.PUT("/:id", UpdateEmployeeHandler(svc)) // Status 200, но records отсутствуют
		employeeGroup.DELETE("/:id", DeleteEmployeeHandler(svc))
	}

	return router
}
