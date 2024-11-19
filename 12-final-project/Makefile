run:
	$(shell grep -v '^#' .env | xargs) go run cmd/my-gram/main.go

build:
	CGO_ENABLED=0 go build -o build/my-gram cmd/my-gram/main.go
