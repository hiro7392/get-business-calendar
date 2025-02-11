FROM golang:1.23 as builder
WORKDIR /workspace

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o get-business-calendar ./cmd/main.go

FROM alpine:latest

RUN apk add --no-cache ca-certificates

COPY --from=builder /workspace/get-business-calendar /workspace/get-business-calendar

EXPOSE 8080

ENV PORT 8080

# アプリケーションの起動
CMD ["/get-business-calendar"]