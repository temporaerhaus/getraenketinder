# syntax=docker/dockerfile:1
FROM golang:latest AS prep
WORKDIR /app
COPY . .
RUN go mod download
RUN go mod verify

FROM prep AS builder
RUN CG0_ENABLED=0 go build -o backend

FROM alpine:latest AS prod
WORKDIR /app
COPY --from=builder /app/backend /backend
ENV GIN_MODE=release
EXPOSE 8888
CMD [ "./backend" ]

FROM prep AS test
CMD ["go", "test", "-v", "./...", "--cover"]