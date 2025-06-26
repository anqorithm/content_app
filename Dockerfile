FROM golang:1.24.4

WORKDIR /usr/src/app

COPY . .
RUN go mod tidy