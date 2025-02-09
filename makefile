# Makefile для создания миграций

# Переменные которые будут использоваться в наших командах (Таргетах)
DB_DSN := "postgres://postgres:yourpassword@localhost:5432/postgres?sslmode=disable"
MIGRATE := migrate -path ./migrations -database $(DB_DSN)

migrate-status:
	migrate -path ./migrations -database $(DB_DSN) version

# Таргет для создания новой миграции
migrate-new:
	migrate create -ext sql -dir ./migrations ${NAME}

# Применение миграций
migrate:
	$(MIGRATE) up

migrate-up-one:
	$(MIGRATE) up 1

# Откат миграций
migrate-down:
	$(MIGRATE) down

migrate-down-one:
	$(MIGRATE) down 1
	
fix-dirty-migrate:
	migrate -path ./migrations -database $(DB_DSN) force ${VERSION}

gen-tasks:
	oapi-codegen -config openapi/.openapi -include-tags tasks -package tasks openapi/openapi.yaml > ./internal/web/tasks/api.gen.go

gen-users:
	oapi-codegen -config openapi/.openapi -include-tags users -package users openapi/openapi.yaml > ./internal/web/users/api.gen.go

lint:
	golangci-lint run --out-format=colored-line-number

# для удобства добавим команду run, которая будет запускать наше приложение
run:
	go run cmd/app/main.go # Теперь при вызове make run мы запустим наш сервер