FROM golang:alpine AS builder

WORKDIR /build
COPY . .

RUN go build .

FROM alpine:3.13.5

WORKDIR /app
EXPOSE 8080
VOLUME ["/app/templates"]
VOLUME ["/app/data"]
COPY --from=builder /build/templates /app/templates
COPY --from=builder /build/wiki /app/wiki

ENTRYPOINT ["/app/wiki"]