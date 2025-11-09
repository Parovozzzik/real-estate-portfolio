FROM golang:1.25-alpine AS builder
WORKDIR /
COPY go.mod go.sum ./
RUN go mod download
COPY . .

ENV GOCACHE=/root/.cache/go-build

RUN --mount=type=cache,target="/root/.cache/go-build" CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go install github.com/pressly/goose/v3/cmd/goose@latest

RUN --mount=type=cache,target="/root/.cache/go-build" CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o real-estate-portfolio ./cmd
#RUN go build -o real-estate-portfolio ./cmd

FROM alpine:latest
WORKDIR /
COPY --from=builder /real-estate-portfolio .
COPY --from=builder /go/bin/goose /usr/local/bin/goose
EXPOSE 8085
CMD ["./real-estate-portfolio"]