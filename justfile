check:
	gofmt -e . > /dev/null
	
install-deps:
	go mod tidy
