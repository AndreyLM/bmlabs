swagger:
	swagger serve ./api/swagger/swagger.yml

run:
	go run ./cmd/server/main.go
	
run-docker:
	cd docker && docker-compose down && docker-compose up -d

build:
	