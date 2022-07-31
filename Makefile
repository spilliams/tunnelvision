.PHONY: test
test:
	go test ./...

.PHONY: build
build:
	go build -o bin/tv cmd/tunnelvision/main.go

.PHONY: install
install:
	go build -o $$GOPATH/bin/tunnelvision cmd/tunnelvision/main.go
