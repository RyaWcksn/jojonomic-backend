FROM golang:alpine as builder

WORKDIR /app 

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o main ./cmd


# Runtime stage
FROM alpine:latest

# Install any required dependencies
RUN apk --no-cache add ca-certificates


# Expose port 9000 for the container
EXPOSE 7002

COPY --from=builder /app/main /usr/bin/


ENTRYPOINT ["main"]