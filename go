> entrypoint.sh
```
#!/bin/sh
set -e

# Run the main application
exec go run .
```

> Dockerfile
```
FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]
```
