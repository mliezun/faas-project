FROM golang:1.16-alpine

WORKDIR /app

COPY . /app

RUN go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o faas main.go

CMD ["/app/faas"]
