FROM golang:1.25-alpine AS builder
WORKDIR /
COPY . .
RUN go install github.com/pressly/goose/v3/cmd/goose@latest
RUN go build -o real-estate-portfolio ./cmd

FROM alpine:latest
WORKDIR /
COPY --from=builder /real-estate-portfolio .
COPY --from=builder /go/bin/goose /usr/local/bin/goose
EXPOSE 8080
CMD ["./real-estate-portfolio"]