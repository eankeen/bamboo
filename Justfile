# build normal binary for GOOS
dev:
	go build .

# build without debugging symbols
prod:
	go build -ldflags '-s -w' .

# update dependencies
update:
	go get -t -u ./...

redep:
	rm go.sum && \
	go get -v -t -d ./...
