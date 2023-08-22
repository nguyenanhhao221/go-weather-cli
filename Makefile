build: 
	@go build -o bin/go-weather-cli

run: build
	@./bin/go-weather-cli

test: 
	@go test -v -cover ./...
