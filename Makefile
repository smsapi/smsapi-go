
lint:
	go vet ./...


tests:
	go test -v ./smsapi


tests-e2e:
	go test -v --tags=e2e ./test/e2e
