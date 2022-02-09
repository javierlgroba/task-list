FROM golang:1.17 AS builder

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o app cmd/server/main.go

FROM alpine:latest as production
COPY --from=builder /app .
CMD ["./app"]