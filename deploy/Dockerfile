# BUILDER Image. Used to download all dependencies, etc
FROM golang:1.19.4 as builder

# Changing root directory
WORKDIR /app

# Copy all files to root directory
COPY . .

# Collect and download dependencies
RUN go mod download

# Building application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main cmd/main.go

# Result image
FROM alpine:3.17

# Changing root directory
WORKDIR /app

# Copy executable file (built application) from builder to root directory
COPY --from=builder /app/main .

# Builded application running
ENTRYPOINT ["/app/main"]

EXPOSE 5554