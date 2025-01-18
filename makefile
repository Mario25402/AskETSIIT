check:
	@echo "Running syntax check"
	gofmt -e ./internal > /dev/null
	
install:
	@echo "Running dependencies installation"
	go mod tidy

test:
	@echo "Running tests"
	go test -v ./internal/... -cover
