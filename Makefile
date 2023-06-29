test:
	@go test

build:
	@go build -o ./bin/wc

run: build
	./bin/wc