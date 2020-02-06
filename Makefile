run:
	go run main.go 

build:
	go build  -v restapi

test:
	go test -v -cover
	

.DEFAULT_GOAL := build



