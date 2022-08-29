FROM golang:1.18-alpine3.16

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download -x

COPY . .


RUN go build -v -o /usr/local/bin/app ./cmd/main.go

EXPOSE 50051

CMD ["app"]