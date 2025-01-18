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
	go test -v ./internal/... -cover
