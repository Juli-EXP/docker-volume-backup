# Build stage
FROM golang:1.21.3-alpine as builder

WORKDIR /app

# Install libraries
COPY go.mod .

COPY go.sum .

RUN go mod download

# Build program
COPY . .

RUN go build -o dvb-server /app/cmd/dvb-server/main.go

#RUN go build -o dvb-cli /app/cmd/dvb-cli/main.go


# Deploy stage 
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/dvb-server /app/dvb-server

#COPY --from=builder /app/dvb-cli /app/dvb-cli

EXPOSE 3000

CMD [ "/app/dvb-server" ]