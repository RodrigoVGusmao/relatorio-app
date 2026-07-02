FROM golang:latest AS builder
WORKDIR /app
#COPY go.mod go.sum ./
#RUN go mod download
COPY . .
RUN go test ./...
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /relatorio-app .

FROM alpine:latest
WORKDIR /root/
RUN apk --no-cache add ca-certificates
COPY --chmod=755 --from=builder /relatorio-app .
CMD ["./relatorio-app"]
