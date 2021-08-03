make:
	go get

build-linux-amd64:
	env GOOS=linux GOARCH=amd64 go build main.go