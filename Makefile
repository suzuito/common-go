
mock:
	./mockgen clogger/a.go
	./mockgen csecret/secret.go
	./mockgen cgcp/memorystore.go
	./mockgen application/a.go

test:
	go test -timeout 30s -coverprofile=coverage.txt -covermode=atomic ./...
	go tool cover -html=coverage.txt -o coverage.html