package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/go-redis/redis"
)

var ErrCacheMiss = errors.New("cache miss")

type RedisClient struct {
	redisClient *redis.Client
}

func (rc *RedisClient) GetByID(employeeID string) (*Employee, error) {

	key := "employee:" + employeeID

	data := rc.redisClient.Get(key)

	if err := data.Err(); err == redis.Nil {
		// not found in cache
		return nil, ErrCacheMiss
	} else if err != nil {
		// another error
		return nil, err
	}

	var employee Employee
	bytesData, err := data.Bytes()
	if err != nil {
		return nil, fmt.Errorf("error when retrieving data from cache: %w", err)
	}

	err = json.Unmarshal(bytesData, &employee)
	if err != nil {
		return nil, fmt.Errorf("error when unmarshalling data from cache: %w", err)
	}

	return &employee, nil
}

func (rc *RedisClient) CreateEmployee(empID, empFirstname, empLastname string) error {

	key := fmt.Sprintf("employee:%s:%s:%s", empID, empFirstname, empLastname)

	fields := map[string]interface{}{
		"employee_id": empID,
		"first_name":  empFirstname,
		"last_name":   empLastname,
	}

	err := rc.redisClient.HMSet(key, fields).Err()
	if err != nil {
		return err // Доделать обработку ошибки
	}

	return nil
}

func (rc *RedisClient) CreatePosition(posID, posName, posEmpID string, posSalary int) error {

	key := fmt.Sprintf("position:%s:%s:%d:%s", posID, posName, posSalary, posEmpID)

	fields := map[string]interface{}{
		"position_id":   posID,
		"position_name": posName,
		"salary":        posSalary,
		"employee_id":   posEmpID,
	}

	err := rc.redisClient.HMSet(key, fields).Err()
	if err != nil {
		return err // Доделать обработку ошибки
	}

	return nil
}

func NewRedisClient(addr, password string, db int) (*RedisClient, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	_, err := client.Ping().Result()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %w", err)
	}

	return &RedisClient{redisClient: client}, nil
}
