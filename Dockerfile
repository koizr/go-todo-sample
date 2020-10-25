FROM golang:1.15.3-alpine AS build

WORKDIR /go/src

COPY . .

RUN apk update && apk add --no-cache git
RUN go build -o go-todo .


FROM alpine

WORKDIR /app

RUN touch .env

COPY --from=build /go/src/go-todo /app/go-todo

CMD ["/app/go-todo"]
