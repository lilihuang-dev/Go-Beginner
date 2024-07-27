build:
	@go build -o bin/Go-Beginner

run: build
	@./bin/Go-Beginner

test:
	@go test -v ./..