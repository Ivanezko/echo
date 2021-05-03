# DEV
FROM golang:1.15 as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/apiserver .

# PROD
FROM alpine:3.9
WORKDIR /app
RUN set -x \
    && apk add --no-cache ca-certificates tzdata \
    && cp /usr/share/zoneinfo/Europe/Kiev /etc/localtime \
    && echo Europe/Kiev > /etc/timezone \
    && apk del tzdata

COPY --from=builder /app/apiserver /app/apiserver

ENV GITHUB-SHA=<GITHUB-SHA>

CMD ["/app/apiserver"]