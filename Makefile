
lint:
	go vet ./...


tests:
	go test -v ./smsapi


tests-e2e:
	go test -v --tags=e2e ./test/e2e


tests-e2e-short-url:
	go test -v test/e2e/smsapi_test.go test/e2e/shorturl_test.go