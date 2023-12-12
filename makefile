migration:
	echo "Hello"

build:
	CGO_ENABLED=0 go build

run:
	go run main.go serve "--http=0.0.0.0:8090"