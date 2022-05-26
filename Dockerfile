FROM golang:1.18-alpine

WORKDIR /sampoerna_notification

COPY . /sampoerna_notification

COPY go.mod /sampoerna_notification
COPY go.sum /sampoerna_notification
RUN go mod download

COPY *.go /sampoerna_notification

RUN go build main.go

CMD go run main.go --mode=stage
