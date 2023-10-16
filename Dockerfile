# Build stage
FROM golang:1.21.3-alpine as build

WORKDIR /app

# Install libraries
COPY go.mod .

RUN go mod download

# Build program
COPY . .

RUN go build -o dvb


# Deploy stage 
FROM alpine:latest

WORKDIR /app

COPY --from=build /app/dbv /app/dvb

EXPOSE 8080

CMD [ "dbv" ]