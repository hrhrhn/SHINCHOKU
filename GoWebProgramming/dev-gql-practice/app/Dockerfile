FROM golang:1.17.3-alpine3.15 as build

WORKDIR /app

RUN apk update --no-cache \
  && apk add --no-cache \
    git \
    gcc \
    musl-dev

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN GOOS=linux GOARCH=amd64 go build -o app main.go

RUN GO111MODULE=off go get github.com/oxequa/realize
RUN GO111MODULE=off go get -tags 'postgres' -u github.com/golang-migrate/migrate/cmd/migrate