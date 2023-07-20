build:
	@go build -o bin/goATM

run: build
	@./bin/goATM

test:
	@go test -v ./...