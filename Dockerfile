FROM golang:1.17.6-alpine3.15
# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

RUN apk --no-cache add gcc g++ make ca-certificates


# Move to working directory /build
WORKDIR /build

# Copy and download dependency using go mod
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the code into the container
COPY . .

# Build the application
RUN go build -a -installsuffix cgo -o main cmd/main.go

# Move to /dist directory as the place for resulting binary folder
WORKDIR /

# Copy binary from build to main folder
RUN cp /build/main .

# Export necessary port
CMD gunicorn --bind 0.0.0.0:$PORT wsgi

# Command to run when starting the container
CMD ["/main"]