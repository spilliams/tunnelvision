.PHONY: build
build:
	go build -o bin/tunnelvision src/cmd/tunnelvision/main.go

.PHONY: install
install:
	go build -o $$GOPATH/bin/tunnelvision src/cmd/tunnelvision/main.go
