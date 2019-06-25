test:
	go test -v -race ./...

bench:
	go test -bench=.
