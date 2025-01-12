check:
	@echo "Running syntax check"
	gofmt -e ./internal > /dev/null
	
install:
	@echo "Running dependencies installation"
	go mod tidy

build:
	@echo "Building project"
	go build ./internal/models/

clean:
	@echo "Cleaning up"
	go clean

test:
	@echo "Running tests"
	go test -v ./internal/models/ -cover

cover:
	@echo "Running coverage"
	go test -coverprofile=coverage.out ./internal/models/
	go tool cover -func=coverage.out
	rm coverage.out