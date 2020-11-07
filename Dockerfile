FROM golang:1.15.2-alpine as builder

RUN mkdir /build
ADD . /build
WORKDIR /build

RUN go mod download
RUN go build -tags docker -o bin/api src/application/main.go

FROM alpine
RUN adduser -S -D -H -h /app appuser
USER appuser

COPY --from=builder /build/bin/api /app/api
COPY --from=builder /build/migrations /migrations

WORKDIR /app
EXPOSE 3001
CMD ["./api"]

