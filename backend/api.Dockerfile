FROM golang:1.22.1-alpine

ENV TZ Asia/Tokyo
ENV GO111MODULE=on

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

RUN go install github.com/cosmtrek/air@latest
CMD ["air", "-c", ".air.toml"]