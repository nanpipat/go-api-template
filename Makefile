dev:
	APP_SERVICE=api go run main.go

# migrate:
# 	docker-compose up -d --build migration

migrate:
	APP_SERVICE=migration go run main.go

test:
	go test ./...
