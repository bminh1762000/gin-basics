build:
	docker-compose build jwt-auth-go

run:
	docker-compose up jwt-auth-go

stop:
	docker-compose stop jwt-auth-go

down:
	docker-compose down

test:
	go test -v ./...

migrate_up:
	migrate -path ./db/migrate -database "postgres://postgres:qwerty@0.0.0.0:5436/postgres?sslmode=disable" up

migrate_down:
	migrate -path ./db/migrate -database "postgres://postgres:qwerty@0.0.0.0:5436/postgres?sslmode=disable" down

swagger:
	swag init -g cmd/main.go
