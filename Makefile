.PHONY: vendor

include bin/build/make/service.mak

# Build release binary.
build:
	go build -race -ldflags="-X 'github.com/alexfalkowski/bezeichner/cmd.Version=latest'" -mod vendor -o bezeichner main.go

# Build test binary.
build-test:
	go test -race -ldflags="-X 'github.com/alexfalkowski/bezeichner/cmd.Version=latest'" -mod vendor -c -tags features -covermode=atomic -o bezeichner -coverpkg=./... github.com/alexfalkowski/bezeichner

sanitize-coverage:
	bin/quality/go/cov

# Get the HTML coverage for go.
html-coverage: sanitize-coverage
	go tool cover -html test/reports/final.cov

# Get the func coverage for go.
func-coverage: sanitize-coverage
	go tool cover -func test/reports/final.cov

# Send coveralls data.
goveralls: sanitize-coverage
	goveralls -coverprofile=test/reports/final.cov -service=circle-ci -repotoken=LOqp8d3hp9KIpdMSCg2tkj54ErsRhdZHB

# Release to docker hub.
docker:
	bin/build/docker/push bezeichner

# Start the environment.
start:
	bin/build/docker/env start

# Stop the environment.
stop:
	bin/build/docker/env stop
