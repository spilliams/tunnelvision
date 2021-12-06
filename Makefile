.PHONY: build
try:
	make build && bin/tv -v file graph tfgraph.dot


.PHONY: build
build:
	go build -o bin/tv src/cmd/tunnelvision/main.go

.PHONY: install
install:
	go build -o $$GOPATH/bin/tunnelvision src/cmd/tunnelvision/main.go
