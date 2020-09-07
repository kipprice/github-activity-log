# iron/go:dev is the alpine image with the go tools added
FROM golang:alpine

# Setup env vars
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app


# install the go app
COPY go.mod . 
COPY go.sum .
RUN go mod download

# Copy source into the container
COPY . . 

# Build it:
RUN go build -o main .

# expose the appropriate port
EXPOSE 8080

# Run the app
CMD ["/app/main"]