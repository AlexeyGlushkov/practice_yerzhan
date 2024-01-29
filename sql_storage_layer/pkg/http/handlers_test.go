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
	err, router := setupTestsResourses()
	if err != nil {
		log.Fatalf("error initializing test resources: %v \n", err)
	}

	testRouter = router
	testServer := httptest.NewServer(testRouter)
	defer testServer.Close()

	exitCode := m.Run()

	if err := teardownTestResourses(); err != nil {
		log.Fatalf("error cleaning test resources: %v \n", err)
	}

	os.Exit(exitCode)
}

func setupTestsResourses() (error, *gin.Engine) {
	db, err := createTestDatabase()
	if err != nil {
		return err, nil
	}
	defer db.Close() // xd

	err = applyMigrations(db)
	if err != nil {
		return err, nil
	}

	// Repo, Cache and Service
	testRepo := repo.NewRepository(db)

	testCache, err := cache.NewRedisClient(testRedisAddr, testRedisPassword, testRedisDB)
	if err != nil {
		return err, nil
	}

	testSvc := service.NewService(*testRepo, *testCache)

	testRouter := SetupRouter(testSvc)

	return nil, testRouter
}

func teardownTestResourses() error {
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
	if err != nil {
		t.Fatalf("Failed to marshal payload to JSON: %v", err)
	}

	req, err := http.NewRequest("POST", "/v1/employee", bytes.NewBuffer(payloadJSON))
	if err != nil {
		t.Fatalf("Failed to create HTTP request: %v", err)
	}

	recorder := httptest.NewRecorder()

	testRouter.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusCreated {
		t.Errorf("Expected status %d; got %d", http.StatusCreated, recorder.Code)
	}

	response := recorder.Body.String()
	expectedResponse := "Employee and Position created successfully"
	if response != expectedResponse {
		t.Errorf("Expected response %s; got %s", expectedResponse, response)
	}
}
