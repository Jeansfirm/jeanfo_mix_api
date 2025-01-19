swag:
	swag init -g cmd/main.go -o docs

test:
	go test ./test/...

test-v:
	go test -v ./test/...