include .env

MIGRATE = -database "mysql://root:root@tcp(localhost:3306)/todo_restful_api?parseTime=true" -path db/migrations

go:
	go run main.go

swag:
	swag init

swag-format:
	swag fmt

migrate-create:
	migrate create -ext sql -dir db/migrations -seq ${name}

migrate-version:
	migrate ${MIGRATE} version

migrate-up:
	migrate ${MIGRATE} up

migrate-down:
	migrate ${MIGRATE} down

migrate-force:
	migrate ${MIGRATE} force ${version}

migrate-drop:
	migrate ${MIGRATE} drop -f