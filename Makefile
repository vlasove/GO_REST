run:
	go run main.go 

build:
	go build 

test:
	go test -v -cover

build2:
	go build  -v restapi

.DEFAULT_GOAL := build2



