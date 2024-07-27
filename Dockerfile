# Stage 1: Build the Go application
FROM golang:1.22 AS build

# Set the working directory inside the container
WORKDIR /srv/grpc

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download and cache dependencies
RUN go mod download

# Copy 
COPY . .

# Build the server
RUN CGO_ENABLED=0 GOOS=linux \
    go build -a -installsuffix cgo \
    -o ./main \
    ./cmd/*


FROM scratch

COPY --from=build /srv/grpc/main /main
    
ENTRYPOINT ["/main"]