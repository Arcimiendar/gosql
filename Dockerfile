FROM golang:1.25.4-trixie

WORKDIR app

COPY go.mod .
RUN go mod download

COPY . .

RUN go build gosql
