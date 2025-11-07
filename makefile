dev:
	cd app && go run main.go
test:
	cd internal/server && go test -v ./...