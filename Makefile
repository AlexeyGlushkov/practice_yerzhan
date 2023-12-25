.SILENT:

.PHONY: lint
lint:
	golangci-lint run

create-migration:
	migrate create -ext sql -dir database/migration/ -seq $(NAME)

migrate-up:
	migrate -path database/migration/ -database "postgresql://postgres:nivea100@localhost:5432/postgres?sslmode=disable" -verbose up

migrate-down:
	migrate -path database/migration/ -database "postgresql://postgres:nivea100@localhost:5432/postgres?sslmode=disable" -verbose down

migrate-fix: 
	migrate -path database/migration/ -database "postgresql://postgres:nivea100@localhost:5432/postgres?sslmode=disable" force $(VERSION)

clean-migration:
	del /Q database\migration\$(FILENAME)
