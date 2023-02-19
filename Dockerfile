FROM golang:1.19-alpine

WORKDIR /usr/app

ENV GOOS=linux \
    GOARCH=amd64

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o /build/book-api.go cmd/main.go

EXPOSE 80

CMD [ "/build/book-api.go" ]