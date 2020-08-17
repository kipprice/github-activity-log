# iron/go:dev is the alpine image with the go tools added
FROM golang:1.12

ENV GO111MODULE=on
WORKDIR /app

COPY go.mod . 
COPY go.sum .

RUN go mod download
COPY . . 

# Build it:
RUN go build -o app

ENTRYPOINT ["./app"]