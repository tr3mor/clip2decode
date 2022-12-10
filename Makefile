build:
	go build -o bin/c2d cmd/clip2decode/main.go

test:
	go test -v ./...