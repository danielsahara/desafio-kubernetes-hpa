FROM golang:1.7-alpine as go-env
RUN mkdir /app

ADD ./src/calcular.go /app

WORKDIR /app

RUN go build calcular.go

ENTRYPOINT ["/app/calcular"]

EXPOSE 80