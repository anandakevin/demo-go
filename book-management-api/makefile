PHONY: vendor
vendor:
	go mod tidy
	go mod vendor

PHONY: server/echo
server/echo:
	go run ./protocol/echo/cmd/main.go

PHONY: server/http
server/http:
	go run ./protocol/http/cmd/main.go

.PHONY: test
test:
	go test -v -race ./...