FROM golang:1.17-alpine as builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -o server cmd/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/server /app

EXPOSE 3000

CMD /app/server