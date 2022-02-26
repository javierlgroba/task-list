FROM golang:1.17-alpine AS builder

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o app cmd/server/main.go

FROM alpine:latest as production
COPY --from=builder /app .
EXPOSE 8080
CMD ["./app"]