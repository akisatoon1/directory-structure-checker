dir-struct-check: go.mod main.go queue.go
	go build

test: dir-struct-check
	../test/test.sh

.PHONY: test