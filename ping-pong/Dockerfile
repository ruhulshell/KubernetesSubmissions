FROM golang:1.26-alpine AS builder

WORKDIR /usr/src/app



COPY go.mod ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o pingpong .

FROM alpine:3.24

COPY --from=builder /usr/src/app/pingpong /app

EXPOSE 8080

CMD ["/app"]