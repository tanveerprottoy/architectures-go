build:
	go build -o bin/app cmd/modular/main.go

run:
	go run cmd/modular/main.go

test-run:
	go test -v -run
