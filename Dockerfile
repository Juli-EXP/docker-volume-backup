# Build stage
FROM golang:1.21.3-alpine as builder

WORKDIR /app

# Install libraries
COPY go.mod .

COPY go.sum .

RUN go mod download

# Build program
COPY . .

RUN go build -o dvb


# Deploy stage 
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/dbv /app/dvb

EXPOSE 3000

CMD [ "dbv" ]