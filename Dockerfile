FROM golang:1.23.2-alpine AS builder

RUN apk add --no-cache gcc musl-dev

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest  

RUN apk --no-cache add ca-certificates sqlite

WORKDIR /root/

COPY --from=builder /app/main .

COPY api_keys.db .

COPY .env .

RUN mkdir -p ./out

EXPOSE 8080

CMD ["./main"]