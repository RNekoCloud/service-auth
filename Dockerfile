FROM golang:1.20.6-alpine3.18

WORKDIR /service-auth

RUN apk update && apk add bash ca-certificates git gcc g++ libc-dev librdkafka-dev pkgconf

COPY go.mod /service-auth

COPY . /service-auth

RUN go mod tidy

RUN go build -tags musl -o /service-auth/bin/main /service-auth/cmd/server.go

EXPOSE 50052 

CMD ["/service-auth/bin/main"]
