BINARY_NAME=aoc24

build:
	@go build -o bin/${BINARY_NAME}

run:
	@go build -o bin/${BINARY_NAME}
	@bin/${BINARY_NAME}