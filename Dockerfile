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

# Install protoc and protoc-gen-go plugin
ARG VERS="25.4"
ARG ARCH="linux-x86_64"
RUN wget https://github.com/protocolbuffers/protobuf/releases/download/v${VERS}/protoc-${VERS}-${ARCH}.zip \
    --output-document=./protoc-${VERS}-${ARCH}.zip && \
    apt update && apt install -y unzip && \
    unzip -o protoc-${VERS}-${ARCH}.zip -d protoc-${VERS}-${ARCH} && \
    mv protoc-${VERS}-${ARCH}/bin/* /usr/local/bin && \
    mv protoc-${VERS}-${ARCH}/include/* /usr/local/include && \
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Build the server
RUN CGO_ENABLED=0 GOOS=linux \
    go build -a -installsuffix cgo \
    -o /go/bin/main \
    ./cmd/*


FROM scratch

COPY --from=build /go/bin/main /main
    
ENTRYPOINT ["/main"]