.PHONY: test

test:
	go test -race -count=1 -p 1 -race -coverpkg=./... -covermode=atomic -coverprofile=coverage.out ./...
