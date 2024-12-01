FROM golang:1.23 AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN go build -o backend -ldflags="-s -w"

FROM alpine:latest
LABEL version="1.0"
LABEL maintainer="Mosin Evgeny <m0s1nevgeny@yandex.ru>"
WORKDIR /app
COPY --from=build /app/backend .
CMD ["./backend"]