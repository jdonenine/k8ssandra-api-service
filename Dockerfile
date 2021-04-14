# Build the manager binary
FROM golang:1.15 as builder

WORKDIR /workspace

# Copy packaging
COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

# Copy source
COPY main.go main.go
COPY pkg/ pkg/

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -o k8ssandra-api-service main.go

# Package
FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY --from=builder /workspace/k8ssandra-api-service .
USER nonroot:nonroot

ENTRYPOINT ["/k8ssandra-api-service"]
