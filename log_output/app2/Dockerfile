FROM golang:1.26-alpine AS builder

WORKDIR /usr/src/app



COPY go.mod ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -v -a -installsuffix cgo -o /usr/local/bin/app ./...

FROM alpine:3.24

RUN adduser -D appuser
USER appuser

COPY --from=builder /usr/local/bin/app /app

EXPOSE 3070

CMD ["/app"]