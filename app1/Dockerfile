# Use a multi-stage build to keep the final image small
FROM golang:1.22.7-alpine AS build-stage

WORKDIR /app

# Install git
RUN apk add --no-cache git

# Set build-time arguments for GitHub credentials
ARG GITHUB_USERNAME
ARG GITHUB_PERSONAL_ACCESS_TOKEN

# Set GOPRIVATE to include your private repositories
ENV GOPRIVATE=github.com/infinid-indonesia

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download Go modules
RUN git config --global url."https://${GITHUB_USERNAME}:${GITHUB_PERSONAL_ACCESS_TOKEN}@github.com/".insteadOf "https://github.com/"
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o ./bin/server ./protocol/http/cmd/main.go

# Deploy the application binary
FROM alpine:3.13
WORKDIR /app
COPY --from=build-stage /app/bin/server /app/server

COPY migrations ./migrations

# Command to run the application
CMD ["./server"]