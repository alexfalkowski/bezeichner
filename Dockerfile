FROM golang:1.21.1-bullseye AS build

ARG version=latest

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN go build -ldflags="-X 'github.com/alexfalkowski/bezeichner/cmd.Version=${version}'" -a -o bezeichner main.go

FROM gcr.io/distroless/base-debian11

WORKDIR /

COPY --from=build /app/bezeichner /bezeichner
ENTRYPOINT ["/bezeichner"]
