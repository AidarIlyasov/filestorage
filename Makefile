build:
	@go build -o ./bin/filestorage
run: build
	@go run filestorage
test:
	@go test ./... -v