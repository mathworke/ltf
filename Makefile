all: deps build


.PHONY: deps
deps:
	go mod init ltf
	go get -d -v github.com/gin-gonic/gin


.PHONY: build
build: deps
	go build -o cmd/ltf.exe cmd/main.go

.PHONY: build_linux
build_linux: deps
	go build -o cmd/ltf cmd/main.go


.PHONY: run
run: 
	go run cmd/main.go


.PHONY: clear
clear:
	rm go.mod
