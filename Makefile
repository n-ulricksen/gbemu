.PHONY: test
test:
	go test -v ./internal/gb/ | tee log.txt
