dev:
	go build .
	./health api
lin:
	GOOS=linux GOARCH=amd64 go build -o health_linux1