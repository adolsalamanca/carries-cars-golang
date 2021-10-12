.PHONY: test

test:
	go test -race -count=1 -p 1 -coverpkg=./... -covermode=atomic -coverprofile=coverage.out ./...
