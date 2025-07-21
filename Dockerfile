# --------------- build stage --------------
FROM golang:1.22 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o notes-api ./cmd/server


# ---------------- runtime stage -----------
FROM debian:bookworm-slim
RUN useradd -m appuser
WORKDIR /app
COPY --from=builder /app/notes-api .
EXPOSE 8080
USER appuser
CMD ["./notes-api"]
