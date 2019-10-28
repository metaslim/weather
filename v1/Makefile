install:
	@dep ensure -v
	
run:
	@go run server/main.go

test:
	@go test -v ./... -race -cover

build:
	go build -v -a -o weather server/main.go