# build executbale binary
FROM golang:alpine as builder
RUN apk update && apk add --no-cache git
WORKDIR /app
COPY . .
RUN go build -o app-service

# 2 build a small image
FROM alpine
WORKDIR /app
COPY --from=builder /app/app-service /app
CMD ["./app-service"]