FROM golang:1.25-alpine AS builder
WORKDIR /
COPY . .
RUN go build -o real-estate-portfolio ./cmd

FROM alpine:latest
WORKDIR /
COPY --from=builder /real-estate-portfolio .
EXPOSE 8080
CMD ["./real-estate-portfolio"]