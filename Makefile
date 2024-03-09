
# For Development
init:
	rm -rf dev.db
	go run migate/migrate.go

dev: 
	CompileDaemon -command="./go-crud"


# For Production
build: 
	go build -o go-crud

prod:
	rm -rf dev.db
	go run migate/migrate.go
	docker-compose up --build

down:
	docker-compose down --rmi all