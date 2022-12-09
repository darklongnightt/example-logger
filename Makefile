start:
	go run examples/logger/main.go
	go run examples/static_logger/main.go

tag:
	git tag v0.1.1

list:
	GOPROXY=proxy.golang.org go list -m github.com/darklongnightt/example-logger