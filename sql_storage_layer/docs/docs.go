// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/employee": {
            "post": {
                "description": "Method for creating a new employee and his position",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Employees"
                ],
                "summary": "Creates an employee and position",
                "parameters": [
                    {
                        "description": "New details",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.CreateEmployeePayload"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Employee and Position created successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid request payload",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/employee/{id}": {
            "get": {
                "description": "Method for retrieving employee details by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Employees"
                ],
                "summary": "Return Employee details by employeeID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "EmployeeID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Employee details",
                        "schema": {
                            "$ref": "#/definitions/main.Employee"
                        }
                    },
                    "404": {
                        "description": "Employee not found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "description": "Method for updating employee details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Employees"
                ],
                "summary": "Updates Employee details",
                "parameters": [
                    {
                        "type": "string",
                        "description": "EmployeeID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "New details",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.UpdateEmployeePayload"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Employee updated successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid request payload",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "Method for deleting employee details by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Employees"
                ],
                "summary": "Deletes Employee details by employeeID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "EmployeeID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Employee deleted successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "main.CreateEmployeePayload": {
            "type": "object",
            "properties": {
                "employee": {
                    "$ref": "#/definitions/main.Employee"
                },
                "position": {
                    "$ref": "#/definitions/main.Position"
                }
            }
        },
        "main.Employee": {
            "type": "object",
            "properties": {
                "employee_id": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                }
            }
        },
        "main.Position": {
            "type": "object",
            "properties": {
                "employee_id": {
                    "type": "string"
                },
                "position_id": {
                    "type": "string"
                },
                "position_name": {
                    "type": "string"
                },
                "salary": {
                    "type": "integer"
                }
            }
        },
        "main.UpdateEmployeePayload": {
            "type": "object",
            "properties": {
                "first_name": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/v1",
	Schemes:          []string{},
	Title:            "Employee Service API",
	Description:      "Application for operations on employee and position",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
