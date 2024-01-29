package main

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

	pkghttp "sql_storage_layer/pkg/http"
	"sql_storage_layer/pkg/models"
	cache "sql_storage_layer/pkg/repo/cache"
	repo "sql_storage_layer/pkg/repo/postgres"
	service "sql_storage_layer/pkg/service"

	"github.com/gin-gonic/gin"
	migrate "github.com/rubenv/sql-migrate"
)

const (
	testHost     = "db"
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

	testRouter := pkghttp.SetupRouter(testSvc)

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
	migration := &migrate.FileMigrationSource{
		Dir: "database/migration",
	}

	num, err := migrate.Exec(db, "postgres", migration, migrate.Up)
	if err != nil {
		return fmt.Errorf("error running migrations: %w \n", err)
	}

	log.Printf("%d migrations applied \n", num)

	return nil
}

func TestCreateEmployeeHandler(t *testing.T) {

	payload := pkghttp.CreateEmployeePayload{
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
