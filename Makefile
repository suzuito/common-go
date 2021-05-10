
mock:
	./mockgen clogger/a.go
	./mockgen csecret/secret.go
	./mockgen cgcp/memorystore.go
	./mockgen application/a.go

test:
	rm -rf /tmp/artifacts && mkdir -p /tmp/artifacts
	go test -timeout 30s -coverprofile=/tmp/artifacts/index.cov ./...
	go tool cover -html=/tmp/artifacts/index.cov -o /tmp/artifacts/index.html