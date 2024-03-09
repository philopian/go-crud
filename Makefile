
# For Development
init:
	rm -rf dev.db
	go run migate/migrate.go

dev: 
	CompileDaemon -command="./go-crud"


# For Production
build:
	rm -rf dist
	rm -rf dev.db
	go run migate/migrate.go
	go build -o dist/go-crud
	GOOS=darwin GOARCH=amd64 go build -o dist/go-crud-macos
	GOOS=linux GOARCH=amd64 go build -o dist/go-crud-linux
	GOOS=windows GOARCH=amd64 go build -o dist/go-crud-windows.exe
	echo "PORT=1234\nDB_FILE_PATH=dev.db" > dist/.env
	cp dev.db dist/dev.db

prod:
	rm -rf dev.db
	go run migate/migrate.go
	docker-compose up --build

down:
	docker-compose down --rmi all