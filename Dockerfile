FROM golang:alpine as builder

RUN mkdir /build
ADD . /build
WORKDIR /build

RUN go mod download
RUN go build -tags docker -o bin/api application/main.go

FROM alpine
RUN adduser -S -D -H -h /app appuser
USER appuser

COPY --from=builder /build/bin/api /app/api
COPY --from=builder /build/migrations /migrations

WORKDIR /app
EXPOSE 8080
CMD ["./api"]

