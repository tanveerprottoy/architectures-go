build:
	go build -o bin/app cmd/modulelayer/main.go

run:
	go run cmd/modulelayer/main.go

test-run:
	go test -v -run

# package wise test
test-user:
	go test -v -cover ./internal/app/modulelayer/module/user

# specific tese function with package path
# go test -run TestMultiply ./

# go test -v <package> -run <TestFunction>
# go test -v -cover --short -race  ./... -run ^TestError*