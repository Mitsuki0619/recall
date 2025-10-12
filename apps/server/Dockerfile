# Build stage
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Install dependencies
RUN apk add --no-cache git ca-certificates

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build binary for AWS Lambda
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-extldflags "-static"' -o bootstrap ./cmd/lambda

# Runtime stage - AWS Lambda compatible
FROM public.ecr.aws/lambda/provided:al2

# Copy binary from builder
COPY --from=builder /app/bootstrap ${LAMBDA_RUNTIME_DIR}

# Set the CMD to your handler
CMD [ "bootstrap" ]