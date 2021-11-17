.PHONY: build
build:
	go build -o bin/tunnelvision cmd/tunnelvision/main.go

.PHONY: install
install:
	go build -o $$GOPATH/bin/tunnelvision cmd/tunnelvision/main.go
