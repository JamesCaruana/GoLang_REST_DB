FROM golang:latest

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

EXPOSE 800 28017

RUN go build

ENTRYPOINT [ "./main" ]