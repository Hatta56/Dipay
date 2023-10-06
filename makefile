install:
	@echo "Pulling all Go dependencies"
	go mod download
	go mod verify
	go mod tidy
	@echo "You can now run 'make build' to compile all packages"

run: 
	go run main.go