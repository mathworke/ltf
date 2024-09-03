all: deps build


.PHONY: deps
deps:
	go mod init ltf
	go get -d -v github.com/gin-gonic/gin


.PHONY: build
build: deps
	go build -o cmd/ltf.exe cmd/ltf/main.go


.PHONY: run
run: 
	go run cmd/ltf/main.go


.PHONY: clear
clear:
	rm go.mod