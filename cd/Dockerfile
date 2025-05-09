FROM golang:1.21 as builder

WORKDIR /app


COPY go.mod ./

RUN go mod download

COPY . .

# 静的リンクでビルド
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o myapp main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/myapp .

ENTRYPOINT ["./myapp"]

EXPOSE 8000

