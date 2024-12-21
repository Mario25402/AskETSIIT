check:
	gofmt -e . > /dev/null
	
install:
	go mod tidy
