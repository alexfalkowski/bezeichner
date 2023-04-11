.PHONY: vendor

include bin/build/make/service.mak

# Build release binary.
build:
	go build -race -ldflags="-X 'github.com/alexfalkowski/bezeichner/cmd.Version=latest'" -mod vendor -o bezeichner main.go

# Build test binary.
build-test:
	go test -race -ldflags="-X 'github.com/alexfalkowski/bezeichner/cmd.Version=latest'" -mod vendor -c -tags features -covermode=atomic -o bezeichner -coverpkg=./... github.com/alexfalkowski/bezeichner

# Release to docker hub.
docker:
	bin/build/docker/push bezeichner
