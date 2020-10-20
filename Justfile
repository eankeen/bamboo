# build normal binary for GOOS
dev:
	go build .

# build without debugging symbols
prod:
	go build -ldflags '-s -w' .

# update dependencies
updep:
	go get -t -u ./...

redep:
	rm go.sum && \
	go get -v -t -d ./...

fmt:
	gofumpt -s -w -extra .

brain:
	just redep
	just updep
	just fmt
