#!make

.default: test

.PHONY: test
test:
	go run cmd/*