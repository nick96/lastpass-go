FROM golang:1.12 AS builder

ENV GO111MODULE=on

WORKDIR /go/src/nick96/lastpass-go
COPY . /go/src/nick96/lastpass-go

RUN go get ./...
CMD go build -v ./...

FROM golang:1.12 AS linter

ENV GO111MODULE=on

WORKDIR /go/src/nick96/lastpass-go
COPY --from=builder /go /go

RUN curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b $(go env GOPATH)/bin v1.17.1

CMD golangci-lint run --enable-all

FROM golang:1.12 AS tester
ENV GO111MODULE=on

WORKDIR /go/src/nick96/lastpass-go
COPY --from=builder /go /go

RUN go get github.com/jstemmer/go-junit-report

CMD go test -v ./... 2>&1 go-junit-report