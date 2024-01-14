package main

import (
	"context"
	"fmt"
)

type Service struct {
	Repo  Repository
	Cache RedisClient
}

func (svc *Service) CreateService(ctx context.Context, emp Employee, pos Position) error {

	tx, err := svc.Repo.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	employeeID, err := svc.Repo.CreateEmployee(ctx, tx, emp.First_name, emp.Last_name)
	if err != nil {
		return fmt.Errorf("repo: failed to create employee: %w", err)
	}

	positionID, err := svc.Repo.CreatePosition(ctx, tx, employeeID, pos.Position_name, pos.Salary)
	if err != nil {
		return fmt.Errorf("repo: failed to create position: %w", err)
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	if err = svc.Cache.CreateEmployee(employeeID, emp.First_name, emp.Last_name); err != nil {
		return fmt.Errorf("cache: failed to create employee: %w", err)
	}

	if err = svc.Cache.CreatePosition(positionID, pos.Position_name, employeeID, pos.Salary); err != nil {
		return fmt.Errorf("cache: failed to create position: %w", err)
	}

	return nil
}

func (svc *Service) GetByIDService(ctx context.Context, employeeID string) (Employee, error) {

	// Пытаемся получить из кэша
	cachedEmp, err := svc.Cache.GetByID(employeeID)
	if err == nil {
		// Запись найдена в кэше, возвращаем её
		return *cachedEmp, nil
	}

	// Запись не найдена в кэше, выполняем поиск в БД
	dbEmp, err := svc.Repo.GetByID(ctx, employeeID)
	if err != nil {
		return Employee{}, fmt.Errorf("GetByID Service error: %w", err)
	}

	// Сохраняем в кэш
	err = svc.Cache.CreateEmployee(dbEmp.Employee_id, dbEmp.First_name, dbEmp.Last_name)
	if err != nil {
		// Обработка ошибки сохранения в кэш, можно проигнорировать или выполнить другие действия
		fmt.Printf("Error caching employee: %v\n", err)
	}

	return dbEmp, nil
}

func (svc *Service) UpdateEmployeeService(ctx context.Context, empID, fName, lName string) error {

	err := svc.Repo.UpdateEmployee(ctx, empID, fName, lName)
	if err != nil {
		return fmt.Errorf("UpdateEmployee Service error: %w", err)
	}

	return nil
}

func (svc *Service) UpdatePositionService(ctx context.Context, posID, posName string, salary int) error {

	err := svc.Repo.UpdatePosition(ctx, posID, posName, salary)
	if err != nil {
		return fmt.Errorf("UpdatePosition Service error: %w", err)
	}

	return nil
}

func (svc *Service) DeleteService(ctx context.Context, employeeID string) error {

	fail := func(err error) error {
		return fmt.Errorf("Delete Service error: %w", err)
	}

	tx, err := svc.Repo.DB.BeginTx(ctx, nil)
	if err != nil {
		return fail(err)
	}
	defer tx.Rollback()

	if err = svc.Repo.Delete(ctx, tx, employeeID); err != nil {
		return fail(err)
	}

	if err = tx.Commit(); err != nil {
		return fail(err)
	}

	return nil
}

func NewService(repo Repository, cache RedisClient) *Service {
	return &Service{
		Repo:  repo,
		Cache: cache,
	}
}
