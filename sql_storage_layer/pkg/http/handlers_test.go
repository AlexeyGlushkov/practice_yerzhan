package http

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"sql_storage_layer/pkg/models"
	cache "sql_storage_layer/pkg/repo/cache"
	repo "sql_storage_layer/pkg/repo/postgres"
	service "sql_storage_layer/pkg/service"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

const (
	testHost     = "localhost"
	testPort     = 5434
	testUser     = "postgres"
	testPassword = "postgres"
	testDbname   = "test"
)

const (
	testRedisAddr     = "localhost:6380"
	testRedisPassword = ""
	testRedisDB       = 0
)

var testRouter *gin.Engine
var testServer *httptest.Server

func TestMain(m *testing.M) {

	exitCode := m.Run()

	os.Exit(exitCode)
}

func setupTestResources() error {
	db, err := createTestDatabase()
	if err != nil {
		return err
	}

	err = applyMigrations(db)
	if err != nil {
		return err
	}

	testRepo := repo.NewRepository(db)

	testCache, err := cache.NewRedisClient(testRedisAddr, testRedisPassword, testRedisDB)
	if err != nil {
		return err
	}

	testSvc := service.NewService(*testRepo, *testCache)

	testRouter = SetupRouter(testSvc)

	testServer = httptest.NewServer(testRouter)

	return nil
}

func teardownTestResources() error {
	testServer.Close()

	return nil
}

func createTestDatabase() (*sql.DB, error) {

	dbConn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		testHost, testPort, testUser, testPassword, testDbname)

	db, err := sql.Open("postgres", dbConn)

	if err != nil {
		return nil, err
	}

	return db, nil
}

func applyMigrations(db *sql.DB) error {
	// SQL-запрос для создания таблицы "employee"
	createEmployeeTable := `
	CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
		CREATE TABLE IF NOT EXISTS employee (
			employee_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
			first_name VARCHAR(50) NOT NULL,
			last_name VARCHAR(50) NOT NULL,
			created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
		);
	`

	// SQL-запрос для создания таблицы "position"
	createPositionTable := `
		CREATE TABLE IF NOT EXISTS position (
			position_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
			position_name VARCHAR(50) NOT NULL,
			salary INTEGER,
			employee_id UUID NOT NULL,
			created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
			CONSTRAINT fk_employee FOREIGN KEY (employee_id) REFERENCES employee (employee_id)
		);
	`

	_, err := db.Exec(createEmployeeTable)
	if err != nil {
		return fmt.Errorf("error creating employee table: %w \n", err)
	}

	_, err = db.Exec(createPositionTable)
	if err != nil {
		return fmt.Errorf("error creating position table: %w \n", err)
	}

	log.Printf("Tables created successfully \n")

	return nil
}

func TestCreateEmployeeHandler(t *testing.T) {

	err := setupTestResources()
	if err != nil {
		t.Fatalf("Failed to set up test resources: %v", err)
	}
	defer func() {
		if err := teardownTestResources(); err != nil {
			t.Fatalf("Failed to tear down test resources: %v", err)
		}
	}()

	payload := CreateEmployeePayload{
		Employee: models.Employee{
			First_name: "John",
			Last_name:  "Doe",
		},
		Position: models.Position{
			Position_name: "Employee",
			Salary:        1000,
		},
	}

	payloadJSON, err := json.Marshal(payload)
	assert.Nil(t, err)

	req, err := http.NewRequest("POST", "/v1/employee/", bytes.NewBuffer(payloadJSON))
	assert.Nil(t, err)

	recorder := httptest.NewRecorder()
	testRouter.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusCreated, recorder.Code)

	expectedJSON := `{"message":"Employee and Position created successfully"}`
	assert.JSONEq(t, expectedJSON, recorder.Body.String())
}

func TestGetEmployeeHandler(t *testing.T) {

	err := setupTestResources()
	if err != nil {
		t.Fatalf("Failed to set up test resources: %v", err)
	}
	defer func() {
		if err := teardownTestResources(); err != nil {
			t.Fatalf("Failed to tear down test resources")
		}
	}()

	employeeID := "c199907e-d190-44ed-bfef-be8a3329e8e2"

	req, err := http.NewRequest("GET", fmt.Sprintf("/v1/employee/%s", employeeID), nil)
	assert.Nil(t, err)

	recorder := httptest.NewRecorder()
	testRouter.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)

	expectedJSON := `{"employee":{"employee_id":"c199907e-d190-44ed-bfef-be8a3329e8e2","first_name":"John","last_name":"Doe"}}`
	assert.JSONEq(t, expectedJSON, recorder.Body.String())
}

func TestUpdateEmployeeHandler(t *testing.T) {

	err := setupTestResources()
	if err != nil {
		t.Fatalf("Failed to setup test resources")
	}
	defer func() {
		if err := teardownTestResources(); err != nil {
			t.Fatalf("Failed to tear down test resources")
		}
	}()

	employeeID := "c199907e-d190-44ed-bfef-be8a3329e8e2"

	payload := UpdateEmployeePayload{
		FirstName: "Doe",
		LastName:  "John",
	}

	payloadJSON, err := json.Marshal(payload)
	assert.Nil(t, err)

	req, err := http.NewRequest("PUT", fmt.Sprintf("/v1/employee/%s", employeeID), bytes.NewBuffer(payloadJSON))
	assert.Nil(t, err)

	recorder := httptest.NewRecorder()
	testRouter.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)

	expectedJSON := `{"message": "Employee updated successfully"}`
	assert.JSONEq(t, expectedJSON, recorder.Body.String())
}

func TestDeleteEmployeeHandler(t *testing.T) {

	err := setupTestResources()
	if err != nil {
		t.Fatalf("Failed to setup test resources")
	}
	defer func() {
		if err := teardownTestResources(); err != nil {
			t.Fatalf("Failed to tear down test resources")
		}
	}()

	employeeID := "c199907e-d190-44ed-bfef-be8a3329e8e2"

	req, err := http.NewRequest("DELETE", fmt.Sprintf("/v1/employee/%s", employeeID), nil)
	assert.Nil(t, err)

	recorder := httptest.NewRecorder()
	testRouter.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)

	expectedJSON := `{"message": "Employee deleted successfully"}`
	assert.JSONEq(t, expectedJSON, recorder.Body.String())
}
