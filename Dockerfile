FROM golang:1.22.4

RUN go version

ENV GOPATH=/
COPY . .
RUN export GOPROXY=direct
RUN go build -o order_service ./cmd/main.go