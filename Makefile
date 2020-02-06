run:
	go run main.go 

build:
	go build  -v restapi

test:
	go test -v -coverprofile=c.out
	go tool cover -html=c.out
	

.DEFAULT_GOAL := build



