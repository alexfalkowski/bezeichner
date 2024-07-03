FROM golang:1.22.5-bullseye AS build

ARG version=latest

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN go build -ldflags="-s -w -X 'github.com/alexfalkowski/bezeichner/cmd.Version=${version}'" -a -o bezeichner main.go

FROM gcr.io/distroless/base-debian12

WORKDIR /

COPY --from=build /app/bezeichner /bezeichner
ENTRYPOINT ["/bezeichner"]
