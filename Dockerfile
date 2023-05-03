FROM golang:1.20-alpine AS build_base

WORKDIR /go/do_ci_cd
COPY go.* ./
RUN go mod download

COPY *.go .
RUN go build *.go

CMD ./main