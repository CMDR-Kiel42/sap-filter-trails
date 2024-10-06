# Build the app
FROM golang:1.23-bookworm as builder
WORKDIR /app
COPY . .
RUN go build -o trails-api /app/main.go

# Then run it
FROM debian:stable-slim
WORKDIR /app
COPY --from=builder /app/trails-api .
COPY ./BoulderTrailHeads.csv /app/
EXPOSE 8080
ENV GIN_MODE=release
CMD ["./trails-api"]
