FROM golang:1.25.1-alpine AS builder

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download

COPY . ./

RUN go build -o server ./app/main.go

FROM alpine:3.20

WORKDIR /app

COPY --from=builder /build/server .

EXPOSE 8080

ENV host="0.0.0.0"
ENV port="8080"
ENV apiVersion="v1"
ENV dbUser="geo"
ENV dbPassword="123"
ENV dbName="wallets"
ENV dbReset="no"

CMD ["./server"]