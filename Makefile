.PHONY: test
test:
	go test -cover -v `glide novendor`

fmt:
	go fmt ./...
