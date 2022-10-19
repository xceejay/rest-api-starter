.PHONY : build run

build:
	@echo "building program..."
	go mod tidy
	go build -o bin/api.out *.go
	ln -sf bin/api.out api.out

build-windows:
	@echo "building program for windows..."
	go mod tidy
	GOOS=windows GOARCH=amd64 go build -o donatepal.exe main.go

run: build
	@echo "executing program..."
	clear
	./api.out
	
