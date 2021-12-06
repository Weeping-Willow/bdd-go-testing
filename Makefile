dep:
	@go get -u github.com/vakenbolt/go-test-report/
	@go get -u github.com/onsi/ginkgo/ginkgo

test:
	@go test ./...

test-fancy:
	@ginkgo tests

test-html:
	@go test ./... -json | go-test-report -g 1