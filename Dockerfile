FROM golang:1.15.3-alpine AS build

WORKDIR /go/src

COPY . .

RUN apk update && apk add --no-cache git
RUN go build -o go-todo .


FROM alpine

WORKDIR /app

COPY --from=build /go/src/go-todo .

RUN addgroup go \
  && adduser -D -G go go \
  && chown -R go:go /app/go-todo

EXPOSE 8080

CMD ["./go-todo"]
