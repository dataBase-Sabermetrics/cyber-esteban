FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

# Build the application specifically for ARM64
RUN GOARCH=arm64 CGO_ENABLED=0 GOOS=linux go build -o /app/main .

EXPOSE 8080

CMD ["/app/main"]
