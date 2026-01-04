# Build stage
FROM golang:1.23-alpine AS builder

RUN apk add --no-cache git make

WORKDIR /app

# Copy dependency files first for better caching
COPY go.mod ./
# RUN go mod download

COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -o bin/whatismyip ./cmd/whatismyip

# Final stage
FROM gcr.io/distroless/static-debian12:nonroot

WORKDIR /app

# Copy binary and web assets
COPY --from=builder /app/bin/whatismyip .
COPY --from=builder /app/web ./web

# Use the nonroot user provided by distroless
USER nonroot:nonroot

EXPOSE 9999

ENTRYPOINT ["./whatismyip"]
