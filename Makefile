
lint:
	go vet ./...


tests:
	go test -v --tags=unit ./smsapi


tests-e2e:
	go test -v --tags=e2e ./smsapi
