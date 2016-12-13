.PHONY: test
test:
	go test -cover `glide novendor`

fmt:
	go fmt ./...
