
init:
	rm -rf dev.db
	go run migate/migrate.go

dev: 
	CompileDaemon -command="./go-crud"

build: 
	go build -o go-crud