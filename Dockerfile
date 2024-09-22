FROM golang:1.23.1-alpine

RUN apk --no-cache --update add gcc g++

WORKDIR /server

COPY ./server/go.mod ./server/go.sum ./

RUN go mod download

COPY ./server .

ENV CGO_ENABLED 1

RUN go build -o ./bin/main .

CMD [ "/server/bin/main" ]
